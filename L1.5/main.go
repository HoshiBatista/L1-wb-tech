package main 

import (
	"fmt"
	"time"
)

func main() {
    const N = 1

    fmt.Printf("Start work for %d seconds...\n", N)

    sendChan := make(chan int)
    done := make(chan struct{})
    timeout := time.After(N * time.Second)

    go func() {
        defer close(sendChan)
        i := 1

        for {
            select {
            case <-done:
                fmt.Println("Sender finished!")
                return
				
            default:
                select {
                case sendChan <- i:
                    fmt.Printf("Sent %d...\n", i)
                    i++
                case <-done:
                    fmt.Println("Sender finished!")
                    return
                }
                time.Sleep(350 * time.Millisecond)
            }
        }
    }()

    go func() {
        for {
            select {
            case data, ok := <-sendChan:

                if !ok {
                    fmt.Println("Receiver finished!")
                    return
                }

                fmt.Printf("Received %d\n", data)

            case <-done:
                fmt.Println("Receiver finished!")
                return
            }
        }
    }()

    <-timeout
    fmt.Println("Time's up! Closing channels...")
    
    close(done) 
    
    time.Sleep(100 * time.Millisecond)

    fmt.Println("Program finished!")
}