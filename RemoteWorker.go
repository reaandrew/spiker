package main

import (
	"fmt"
	"io"
	"log"
)

type RemoteWorkerClient struct{
	id int32
	worker LoadrService_ConnectServer
}

func (remoteWorker *RemoteWorkerClient) Run(specification *TestSpecification){
	go func(worker LoadrService_ConnectServer){
		log.Println("Send worker request inside ControlRequest")

		worker.Send(&ControlRequest{
			Spec: specification,
		})
	}(remoteWorker.worker)
}

func (remoteWorker *RemoteWorkerClient) Wait(){
	for {
		_, err := remoteWorker.worker.Recv()
		if err == io.EOF || err != nil{
			fmt.Println("Worker disconnected")
			break
		}
	}
}

func NewRemoteWorker(worker LoadrService_ConnectServer, id int32) *RemoteWorkerClient {
	return &RemoteWorkerClient{worker: worker, id: id}
}
