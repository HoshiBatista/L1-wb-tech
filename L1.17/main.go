package main

import (
	"cmp"
	"fmt"
)

func BinarySearch[T cmp.Ordered](arr []T, target T) int {
	l := 0
	r := len(arr) - 1

	for l <= r {
		mid := l + (r-l)/2

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func main() {
	intArr := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("Поиск 7 в %v: индекс %d\n", intArr, BinarySearch(intArr, 7))

	strArr := []string{"apple", "banana", "cherry", "date", "e"}
	fmt.Printf("Поиск 'e' в %v: индекс %d\n", strArr, BinarySearch(strArr, "e"))

	floatArr := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Printf("Поиск 1.1 в %v: индекс %d\n", floatArr, BinarySearch(floatArr, 1.1))
}
