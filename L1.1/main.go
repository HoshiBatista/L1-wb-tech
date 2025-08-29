package main

import (
	"fmt"
)

type Human struct {
	Name  	 string
	Surname  string
	Age		 int
}

func (h *Human) Hello() {
	fmt.Printf("Hi! I'm %s %s\n", h.Name, h.Surname)
}

func (h *Human) Old() {
	fmt.Printf("I'm %d years old\n", h.Age)
}

type Action struct {
	Counter int
	Human
}

func NewAction(cnt int, h Human) *Action {
	return &Action{
		Counter: cnt,
		Human: h,
	}
}

func main() {
	var act Action = *NewAction(0, Human{"Roman", "Moroz", 20}) 

	act.Hello()
	act.Old()
}