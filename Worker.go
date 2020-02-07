package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"io"
	"log"
	"time"
)

type Worker struct {
	address string
}

func NewWorker(address string) *Worker {
	return &Worker{address: address}
}

func (worker *Worker) connect(ctx context.Context) (LoadrService_ConnectClient, LoadrServiceClient){
	var client LoadrService_ConnectClient
	var c LoadrServiceClient

	connected := false

	for {
		if connected {
			break
		}

		log.Println("Worker connecting to server")
		fmt.Println("Dialing the server...")
		var kacp = keepalive.ClientParameters{
			Time:                10 * time.Second, // send pings every 10 seconds if there is no activity
			Timeout:             time.Second,      // wait 1 second for ping ack before considering the connection dead
			PermitWithoutStream: true,             // send pings even without active streams
		}
		conn, err := grpc.Dial(worker.address, grpc.WithInsecure(), grpc.WithKeepaliveParams(kacp))
		if err != nil {
			conn.Close()
			time.Sleep(1*time.Second)
		}else{
			c = NewLoadrServiceClient(conn)
			client, err = c.Connect(ctx)
			if err != nil {
				fmt.Println("Sleeping")
				time.Sleep(1*time.Second)
			}else{
				connected = true
			}
		}
	}

	return client, c
}

func (worker *Worker) Start() {
	//This will connect to the server and then run the necessary commands
	//Even if it is standalone, this will still use the distributed method
	// but with a single worker node.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, c := worker.connect(ctx)

	for {
		controlRequest, err := client.Recv()
		if err == io.EOF || err != nil {
			fmt.Println("Server disconnected, reconnecting")
			client, c = worker.connect(ctx)
		}else {
			service := NewService()
			result, err := service.Run(controlRequest.Spec)
			if err != nil {
				panic(err)
			}
			fmt.Println("Sending result")
			c.SendTestResult(ctx, &SendTestResultRequest{
				Result: &result,
			})
		}
	}
}
