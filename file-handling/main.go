package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "Hello There, I am Krishna"
	file, err := os.Create("./filename.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	file.Close()
}
