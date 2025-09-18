package main

import (
	"fmt"
	"sync"
)

func main() {
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	wg.Add(1)

	go func() {
		defer close(input)

		for _, n := range numbers {
			input <- n
		}

		wg.Done()
	}()

	wg.Add(1)

	go func() {
		defer close(output)

		for n := range input {
			output <- n * 2
		}

		wg.Done()
	}()

	wg.Add(1)

	go func() {
		for result := range output {
			fmt.Println(result)
		}

		wg.Done()
	}()

	wg.Wait()
}
