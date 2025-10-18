package main

import (
	"fmt"
)

func QuickSort(arr []int) []int {
    if len(arr) < 2 {
        return arr
    }
    
    left, right := 0, len(arr)-1
    
    pivot := len(arr) / 2
    
    arr[pivot], arr[right] = arr[right], arr[pivot]
    
    for i := range arr {
        if arr[i] < arr[right] {
            arr[left], arr[i] = arr[i], arr[left]
            left++
        }
    }
    
    arr[left], arr[right] = arr[right], arr[left]
    
    QuickSort(arr[:left])
    QuickSort(arr[left+1:])
    
    return arr
}

func main() {
    testArrays := [][]int{
        {64, 34, 25, 12, 22, 11, 90},
        {5, 2, 3, 1, 4},
        {1},
        {},
        {3, 3, 3, 3},
        {9, 8, 7, 6, 5, 4, 3, 2, 1},
    }
    
    fmt.Println("=== Реализация QuickSort ===")
    for i, arr := range testArrays {
        original := make([]int, len(arr))

        copy(original, arr)
        result := QuickSort(arr)
		
        fmt.Printf("Test %d: %v -> %v\n", i+1, original, result)
    }
}