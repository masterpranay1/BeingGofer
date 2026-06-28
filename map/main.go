package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["Go"] = "Golang"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
