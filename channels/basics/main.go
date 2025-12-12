package main

import (
	"fmt"
	"time"
)

func main() {
	defer fmt.Println("Main goroutine exited.")
	defer fmt.Println("Main function completed")
	defer fmt.Println("checking defer")
	message := make(chan string)
	message2 := make(chan int)
	message3 := make(chan int)
	go func() {
		time.Sleep(time.Second * 2)
		message <- "Hello from goroutine!"
		message <- "one more time"
		message2 <- 1
		message3 <- 3
	}()
	msg := <-message
	fmt.Println(msg)

	msg = <-message
	fmt.Println(msg)

	time.Sleep(time.Second * 2)
	//close(message2)

	msg2 := <-message2
	fmt.Println(msg2)

	close(message2)
	//msg := <-message

	fmt.Println(msg)
	close(message)
	//close(message)
	//close(message2)
	close(message3)

	msg3 := <-message3
	fmt.Println(msg3)
}
