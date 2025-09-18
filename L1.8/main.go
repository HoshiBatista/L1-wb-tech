package main

import (
	"fmt"
	"strconv"
)

func setBit(num int64, i uint, bit int) int64 {
	if bit == 1 {
		return num | (1 << i)
	} else {
		return num &^ (1 << i)
	}
}

func main() {
	var num int64 = 7
	i := uint(1)

	fmt.Printf("Исходное число: %d (%s)\n", num, strconv.FormatInt(num, 2))

	result := setBit(num, i, 0)
	fmt.Printf("После установки %d-го бита в 0: %d (%s)\n", i, result, strconv.FormatInt(result, 2))

	result = setBit(num, i, 1)
	fmt.Printf("После установки %d-го бита в 1: %d (%s)\n", i, result, strconv.FormatInt(result, 2))
}
