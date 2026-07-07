package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://example.com"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	response, err := http.Get(url)
	checkError(err)

	// type of the the response
	fmt.Printf("Type of Response is %T\n", response)
	// OUTPUT : Type of Response is *http.Response

	// this is very important to close manually
	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	checkError(err)
	content := string(databytes)
	fmt.Println(content)
}
