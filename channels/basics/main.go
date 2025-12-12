package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer fmt.Println("Main goroutine exited.")
	defer fmt.Println("Main function completed")
	defer fmt.Println("checking defer")

	message := make(chan string)
	message2 := make(chan int)
	message3 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)

		message <- "Hello from goroutine!"
		message <- "one more time"
		message2 <- 1
		message3 <- 3

		close(message)
		close(message2)
		close(message3)
	}()

	msg := <-message
	fmt.Println(msg)

	msg = <-message
	fmt.Println(msg)

	msg2 := <-message2
	fmt.Println(msg2)

	msg3 := <-message3
	fmt.Println(msg3)

	wg.Wait()
}
