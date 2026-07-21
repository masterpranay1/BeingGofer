package osMethods

import (
	"fmt"
	"os"
)

func RFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(data)
}
