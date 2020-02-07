package main

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"log"
	"net"
	"sync/atomic"
	"time"
)

type RemoteController struct {
	workerIds int32
	workers map[int32]*RemoteWorker
	server  *grpc.Server
	waitFor int
	results int
	onReadySubscribers []func()
	onCompleteSubscribers []func()
}

func (r *RemoteController) OnReady(subscriber func()){
	r.onReadySubscribers = append(r.onReadySubscribers, subscriber)
}

func (r *RemoteController) publishOnReady(){
	fmt.Println("Publishing on ready")
	for _, subscriber := range r.onReadySubscribers{
		fmt.Println("Publishing on ready to subscriber")
		subscriber()
	}
}

func (r *RemoteController) OnComplete(subscriber func()){
	r.onCompleteSubscribers = append(r.onCompleteSubscribers, subscriber)
}

func (r *RemoteController) publishOnComplete(){
	fmt.Println("Publishing OnComplete")
	for _, subscriber := range r.onCompleteSubscribers{
		subscriber()
	}
	r.server.Stop()
}

func (r *RemoteController) OnWorkerConnected(){
	fmt.Println("OnWorkerConnected",r.waitFor, len(r.workers))
	if r.waitFor > 0 && len(r.workers) == r.waitFor{
		r.publishOnReady()
	}
}

func (r *RemoteController) OnTestResult(_ *TestResult){
	r.results++
	fmt.Println("OnTestResult", r.results, len(r.workers))
	if r.results == len(r.workers){
		fmt.Println("Publishing complete...")
		r.publishOnComplete()
	}
}

func (r *RemoteController) SendTestResult(_ context.Context, result *SendTestResultRequest) (*SendTestResultResponse, error) {
	r.OnTestResult(result.Result)
	return &SendTestResultResponse{}, nil
}

func (r *RemoteController) Run(spec *TestSpecification) (TestResult, error) {
	fmt.Println("Sending the spec to each worker", len(r.workers))
	for index, worker := range r.workers {
		log.Println("Sending the spec to each worker", index+1)
		worker.Run(spec)
	}
	return TestResult{}, nil
}

func (r *RemoteController) addWorker(worker *RemoteWorker) {
	r.workers[worker.id] = worker
}

func (r *RemoteController) Connect(worker LoadrService_ConnectServer) error {
	id := atomic.AddInt32(&r.workerIds, 1)
	remoteWorker := NewRemoteWorker(worker, id)
	r.addWorker(remoteWorker)
	defer func(){
		r.removeWorker(remoteWorker)
		fmt.Println("Number of connected workers", len(r.workers))
	}()
	fmt.Println("Number of connected workers", len(r.workers))
	r.OnWorkerConnected()
	remoteWorker.Wait()

	return nil
}

func (r *RemoteController) Start(bindAddress string) {
	lis, err := net.Listen("tcp", bindAddress)
	if err != nil {
		log.Fatalf("Failed to listen on port: %v", err)
	}
	fmt.Println("Server Listening...")
	if err := r.server.Serve(lis); err != nil {
		log.Fatal(errors.Wrap(err, "Failed to start server!"))
	}
}

func (r *RemoteController) removeWorker(worker *RemoteWorker) {
	delete(r.workers, worker.id)
}

func NewRemoteController(waitFor int) *RemoteController {
	controller := &RemoteController{
		waitFor: waitFor,
		onReadySubscribers: []func(){},
		onCompleteSubscribers: []func(){},
		workers: map[int32]*RemoteWorker{},
	}

	var kaep = keepalive.EnforcementPolicy{
		MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
		PermitWithoutStream: true,            // Allow pings even when there are no active streams
	}
	server := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep))
	RegisterLoadrServiceServer(server, controller)
	controller.server = server
	return controller
}
