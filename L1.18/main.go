package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	value int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

func (c *AtomicCounter) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.value)
}

func main() {
	counter := NewAtomicCounter()
	var wg sync.WaitGroup

	const goroutines = 1000
	const incrementsPerGoroutine = 100

	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Printf("Итоговое значение атомарного счетчика: %d\n", counter.Value())
	fmt.Printf("Ожидаемое значение: %d\n", goroutines*incrementsPerGoroutine)
}