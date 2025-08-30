package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("Worker %d, chanel close, finish job ... \n", id)
				return
			}

			fmt.Printf("Worker %d start job %d ... \n", id, job)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d finish job %d ... \n", id, job)

		case <-ctx.Done():
			fmt.Printf("Worker %d get cancel signal ... Finish job ... \n", id)
			return
		}
	}
}

func main() {
	const numDefaultWorkers = 5

	numWorkers := flag.Int("workers", numDefaultWorkers, "Count of workers")
	flag.Parse()

	const nubJobs = 50

	jobs := make(chan int, nubJobs)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	for w := 1; w < *numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, &wg)
	}

	go func() {
		signalChanel := make(chan os.Signal, 1)
		signal.Notify(signalChanel, syscall.SIGINT, syscall.SIGTERM)
		<-signalChanel

		fmt.Printf("\n Get canc0el signal! Stoping ...  \n")
		cancel()
	}()

	go func() {
		defer close(jobs)

		for i := 1; ; i++ {
			select {
			case jobs <- i:
				fmt.Println("Successfully submitted task ...")

			case <-ctx.Done():
				fmt.Println("Stopping sending tasks ...")
				return
			}
		}
	}()

	wg.Wait()

	fmt.Println("Program finished")
}
