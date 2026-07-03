package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

// Very basic method
func (u User) GetStatus() bool {
	return u.Status
}

func main() {
	fmt.Println("Methods in Go")
	user := User{
		Name:   "Pranay",
		Email:  "pranay@gmail.com",
		Status: true,
		Age:    24,
	}

	fmt.Println(user.GetStatus())
}
