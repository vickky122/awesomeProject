package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	user1 := user{
		firstName: "John",
		lastName:  "Doe",
		birthDate: "1990-01-01",
		createdAt: time.Now(),
	}
	user2 := user{
		firstName: "vikrant",
		lastName:  "yadav",
		birthDate: "1990-01-01",
	}
	fmt.Println(user1, user2)
}
