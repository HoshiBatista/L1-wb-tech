package main

import "fmt"

func swap(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {
	x, y := 5, 7
	fmt.Printf("До обмена: \n x = %d, y = %d \n", x, y)

	swap(&x, &y)

	fmt.Printf("После обмена: \n x = %d, y = %d \n", x, y)
}