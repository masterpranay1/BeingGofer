package main

import (
	"fmt"
	osMethods "hello/os-methods"
	"os"
)

func main() {
	// content := "Hello There, I am Krishna"
	// file, err := os.Create("./filename.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// length, err := io.WriteString(file, content)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(length)
	// defer file.Close()

	// readFile("./filename.txt")

	osMethods.RFile("./filename.txt")
}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	// databyte will not be a string type

	if err != nil {
		panic(err)
	}

	fmt.Println("Text inside file is", string(databyte))
}
