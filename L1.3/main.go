package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("\nWorker %d start job: %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("\nWorker %d finish job: %d\n", id, job)
	}
}

func main() {
	const numDefaultWorkers = 5

	numWorkers := flag.Int("workers", numDefaultWorkers, "Count of workers")
	flag.Parse()

	const numJobs = 10

	jobs := make(chan int, numJobs)

	var wg sync.WaitGroup

	fmt.Printf("Start %d workers ...\n", *numWorkers)

	for w := 1; w <= *numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	for j := 1; j < numJobs; j++ {
		jobs <- j
	}

	fmt.Printf("Close jobs ...")

	close(jobs)

	wg.Wait()

	fmt.Println("\nAll workers finish!")
}
