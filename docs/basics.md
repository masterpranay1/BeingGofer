# GoLang Course

## Introduction

- Widely used in cloud infrastructure
- Go - Next Generation Language
- Compiled Language
- Don't judge the syntax 😂

**Installation and first "go"**

- Install go on your pc using simple download and installation
- Start with this basic code:

```go
package main

import "fmt"

func main() {
	fmt.Println("My first go code")
}
```

**Questions?**

1. Should package name must be same with the file name? What is a package after all?
2. Is the main function "main" compulsory?

- Go will automatically import the file, no need to self-write the import statement

**Concern:** How it works in case of duplicates?

**Running the go code**

- `go run main.go` will the directly run the code.
- `go build main.go` will build a executable `main` 

**Cool Commands**

- `go env GOPATH` shows the path of the go where it is installed; Be careful of upper and lower case.

**Lexer**

- We don't have to put semicolons everywhere, Lexer does this job for us, when we run the go code.

**Types**

- Strings, Bool, Integer, Floating, Complex
- Integer extra types: uint8, uint64, int8, int64, uintptr
- Floating extra types: float32, float64
- Some advanced types: Array, Slices, Maps, Structs, Pointers
- Many other types are also there
