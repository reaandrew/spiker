package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	waitFor := 3
	serverAddress := "127.0.0.1:41521"

	if os.Args[1] == "worker" {
		worker := NewWorker(serverAddress)
		time.Sleep(1 * time.Second)
		worker.Start()
	}

	if os.Args[1] == "server" {
		wg := sync.WaitGroup{}
		wg.Add(1)

		controller := NewRemoteController(waitFor)
		controller.OnReady(func() {
			fmt.Println("Running test")
			controller.Run(&TestSpecification{})
		})
		controller.OnComplete(func() {
			fmt.Println("Completed!!")
			wg.Done()
		})

		time.Sleep(2 * time.Second)
		go controller.Start(serverAddress)

		wg.Wait()
	}
}
