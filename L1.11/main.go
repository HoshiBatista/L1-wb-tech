package main

import (
	"fmt"
	"sort"
)

func sliceIntersection(sl1, sl2 []int) []int {

	sort.Ints(sl1)
	sort.Ints(sl2)

	var result []int
	i, j := 0, 0

	for i < len(sl1) && j < len(sl2) {

		if i > 0 && sl1[i] == sl1[i-1] {
			i++
			continue
		}

		if j > 0 && sl2[j] == sl2[j-1] {
			j++
			continue
		}

		if sl1[i] == sl2[j] {
			result = append(result, sl1[i])
			i++
			j++
		} else if sl1[i] < sl2[j] {
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}
	
	fmt.Print("Пересечение: ")
	fmt.Println(sliceIntersection(A, B))
}