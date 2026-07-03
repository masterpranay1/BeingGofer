package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	// private field as first letter is small
	address Address
}

type Address struct {
	City    string
	State   string
	Country string
}

func main() {
	user := User{
		Name:   "Pranay",
		Email:  "pranay@gmail.com",
		Status: true,
		Age:    24,
		address: Address{
			City:    "Mumbai",
			State:   "Maharashtra",
			Country: "India",
		},
	}
	fmt.Println(user)
	fmt.Printf("User details are: %+v\n", user)
	fmt.Printf("Name is %v and email is %v\n", user.Name, user.Email)
	fmt.Printf("Address is %v\n", user.address)
}
