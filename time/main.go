package main

import (
	"fmt"
	"time"
)

func main() {
	presentTime := time.Now()
	fmt.Println(presentTime)

	// Formatting the time
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
}
