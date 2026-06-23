## Types

```jsx
package main

import "fmt"

func main() {
	var name string = "John"
	fmt.Println(name)
}
```

- You cannot declare a variable and don’t use it. We will get an error.

**Printf**

```go
fmt.Printf("Variable is of type %t \n", name)
```

**Question :** What’s the difference in Printf and Println ?

**float32 and float64**

```go
var amount float32 = 45.123423424234
```

This will be trimmed to lesser decimal points automatically, and similarly float64 will also be trimmed to lesser decimal points, but off-course greater than float32

**Inference**

- You can ignore the type
    
    `var name = "Pranay"` 
    

**No Var Style** 

- `userCnt := 10`
- Can only be used inside a function and not in global scope

**Default Values** 

- Variables declared without a corresponding initialization are *zero-valued*. For example, the zero value for an `int` is `0`.

**Multiple Declarations**

- `var b, c int = 1,2`