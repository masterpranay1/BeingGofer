package main

import "fmt"

func main() {
	var name string = "John"
	fmt.Println(name)

	var age int = 20
	fmt.Println(age)

	var isStudent bool = true
	fmt.Println(isStudent)

	var pi float32 = 3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679
	fmt.Println(pi)
	// only 6 decimal places are shown

	// Alias
	aliasName := "John"
	fmt.Println(aliasName)
}
