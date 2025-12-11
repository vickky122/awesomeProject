package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello")
	time.Sleep(5 * time.Second)
	fmt.Println("This is Vikrant")
}

func sayBye() {
	fmt.Println("I am leaving")
	time.Sleep(2 * time.Second)
	fmt.Println("Bye Bye")
}

func main() {
	fmt.Println("starting goroutine")
	go sayHello()
	go sayBye()

	time.Sleep(6 * time.Second)
}
