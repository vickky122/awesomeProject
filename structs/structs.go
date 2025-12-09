package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func NewUser(first, last, dob string) (*User, error) {
	if first == "" || last == "" || dob == "" {
		return nil, errors.New("invalid user data")
	}
	return &User{
		firstName: first,
		lastName:  last,
		birthDate: dob,
		createdAt: time.Now(),
	}, nil
}

func (u User) FullName() string {
	return u.firstName + " " + u.lastName
}

type Admin struct {
	User
	Password string
}

func NewAdmin(first, last, dob, password string) (*Admin, error) {
	user, err := NewUser(first, last, dob)
	if err != nil {
		return nil, err
	}

	if password == "" {
		return nil, errors.New("password required")
	}

	return &Admin{
		User:     *user,
		Password: password,
	}, nil
}

func (a Admin) ShowAdminDetails() {
	fmt.Println("Admin Name:", a.FullName())
	fmt.Println("Birth Date:", a.birthDate)
	fmt.Println("Account Created:", a.createdAt)
	fmt.Println("Admin Password:", a.Password)
}

func main() {
	admin, err := NewAdmin("Vikrant", "Yadav", "1777-07-07", "supersecret123")
	if err != nil {
		fmt.Println(err)
		return
	}

	admin.ShowAdminDetails()
}
