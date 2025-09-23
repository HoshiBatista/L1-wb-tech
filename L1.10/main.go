package main

import (
	"fmt"
	"sort"
)

func groupTemp(temps []float64) map[int][]float64 {
	hashMap := make(map[int][]float64)

	for _, el := range temps {
		key := int(el/10) * 10

		hashMap[key] = append(hashMap[key], el)
	}

	return hashMap
}

func formatOutput(hashMap map[int][]float64) {
	var keys []int

	for k := range hashMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for idx, key := range keys {
		fmt.Printf("%d:{", key)

		for i, temp := range hashMap[key] {
			if i > 0 {
				fmt.Print(", ")
			}

			fmt.Printf("%.1f", temp)
		}

		if idx != len(keys)-1 {
			fmt.Print("}, ")
		} else {
			fmt.Println("}")
		}
	}
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	formatOutput(groupTemp(temperatures))
}
