package main

import (
	"fmt"
	"sync"
)

func squareArrElements(arr []int, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Goroutine %d start\n", id)

	for idx, el := range arr {
		arr[idx] = el * el
	}

	fmt.Printf("Goroutine %d finish\n", id)
}

func main() {
	arr1 := []int{2, 4, 6, 8, 10}
	arr2 := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	wg.Add(2)

	go squareArrElements(arr1, 1, &wg)
	go squareArrElements(arr2, 2, &wg)

	wg.Wait()

	fmt.Println(arr1)
	fmt.Println(arr2)
}
