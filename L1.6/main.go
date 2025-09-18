package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func exitByCondition() {
	fmt.Println("=== Выход по условию ===")

	var wg sync.WaitGroup
	var stopFlag int32

	wg.Add(1)
	go func() {
		defer wg.Done()

		for atomic.LoadInt32(&stopFlag) == 0 {
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)
		}
		fmt.Println("Горутина завершена по условию ...")
	}()

	time.Sleep(2 * time.Second)
	atomic.StoreInt32(&stopFlag, 1)

	wg.Wait()

	fmt.Println()
}

func exitByChannel() {
	fmt.Println("=== Через канал уведомления ===")

	stopChan := make(chan struct{})
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-stopChan:
				fmt.Println("Горутина завершена по сигналу канала ...")
				return
			default:
				fmt.Println("Горутина работает ...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	close(stopChan)
	<-done

	fmt.Println()
}

func exitByContext() {
	fmt.Println("=== Через контекст ===")

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина завершена по сигналу контекста ...")
				return
			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	cancel()
	<-done

	fmt.Println()
}

func exitByRuntime() {
	fmt.Println("=== Использование runtime.Goexit() ===")

	done := make(chan struct{})

	go func() {
		defer close(done)
		defer fmt.Println("Горутина завершена через runtime.Goexit() ...")

		for {
			fmt.Println("Горутина работает...")
			time.Sleep(500 * time.Millisecond)

			if time.Now().Unix()%5 == 0 {
				runtime.Goexit()
			}
		}
	}()

	<-done

	fmt.Println()
}

func exitByChannelClose() {
	fmt.Println("=== Через закрытие канала данных ===")

	dataChan := make(chan int)
	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			select {
			case data, ok := <-dataChan:
				if !ok {
					fmt.Println("Горутина завершена после закрытия канала ...")
					return
				}

				fmt.Printf("Получены данные: %d\n", data)

			default:
				fmt.Println("Горутина работает...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second)

	close(dataChan)
	<-done

	fmt.Println()
}

func main() {
	exitByCondition()    // 1. Выход по условию
	exitByChannel()      // 2. Через канал уведомления
	exitByContext()      // 3. Через контекст
	exitByRuntime()      // 4. Использование runtime.Goexit()
	exitByChannelClose() // 5. Через закрытие канала данных
}
