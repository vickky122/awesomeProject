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

func (u user) fullName() string {
	return u.firstName + " " + u.lastName
}
func (u user) updateFirstName(newFirstName string) {
	u.firstName = newFirstName
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
		createdAt: time.Now(),
	}
	user3 := &user{
		firstName: "vikrant",
		lastName:  "yadav",
		birthDate: "1990-01-01",
		createdAt: time.Now(),
	}
	fmt.Println(user1)
	fmt.Println(user2)
	fmt.Println(*user3)

	user2.fullName()
	user3.updateFirstName("vicky")

	fmt.Println(user2)
	fmt.Println(*user3)
}
