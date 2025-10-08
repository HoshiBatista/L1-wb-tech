package main

import "fmt"

func DetermineType(v interface{}) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	case chan bool:
		return "chan bool"
	case chan interface{}:
		return "chan interface{}"
	default:
		return fmt.Sprintf("unknown type: %T", v)
	}
}

func main() {
	fmt.Println("\nДемонстрация работы в runtime:")

	var (
		num     = 100
		text    = "world"
		flag    = false
		intCh   = make(chan int)
		strCh   = make(chan string)
	)
	
	fmt.Printf("num (%v) имеет тип: %s\n", num, DetermineType(num))
	fmt.Printf("text (%v) имеет тип: %s\n", text, DetermineType(text))
	fmt.Printf("flag (%v) имеет тип: %s\n", flag, DetermineType(flag))
	fmt.Printf("intCh (%v) имеет тип: %s\n", intCh, DetermineType(intCh))
	fmt.Printf("strCh (%v) имеет тип: %s\n", strCh, DetermineType(strCh))
}