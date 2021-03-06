## GOPL

```go 
// Package comment to describe the package before declaring package
package main


import (
  "fmt"
)
```

## Chapter 1
### Chapter 1.1
- `go build helloworld.go` to build the binary to run using `./helloworld` in the future
- Alternatively run `go run helloworld.go` to run the file without compiling
- `gofmt` should be run when save because go takes a strong stance on code formatting
- Program will not compile if there are missing imports or if there are uneccesarry ones.

### Chapter 1.2 
- 
- `:=` symbol is part of a short variable declaration. Giving them appropriate types based on initializer values
- `i++` is +1, `i--` is minus 1. 

- for loop comes in the form of: 
```go
for initialization; condition; post {
  // statement 
  // parentheses are never used around the three components of for loop
  // initialization is optional and usually a short statement 
  // loop runs when condition is true
}
```
```go
// A more traditional for loop
for condition {
  // statement
  }
```

- `range` == `enumerate` in python, providing index and an element
```go
  // ch1/echo2
  for _, arg := range os.Args[1:] {

  }
```

- Use one of the first 2 forms as others are redundant. Explicit  initialization when initial value matters and implicit when initial value doesn't 
```go
s := ""
var s string
var s = ""
var s string = ""
```
The first form can only be used in a function.
The second form relies on default initialization to zero value. 
The third form is rarely used when declaring multiple variables

- An easier way to iterate through and concatenatingo
```go
strings.Join(os.Args[1:], "") 
```

- os.Args[0] is the name of the program

### Chapter 1.3

```go
func main() {
counts := make(map[string]int)
input := bufio.NewScanner(os.Stdin)

for input.Scan() {
  counts[input.Text()]++
}

for line, n := range counts {
  if n > 1 {
    fmt.Printf("%d\t%s\n",n,line)
    }
  }
}
```

- Parentheses are not used for if statements but braces are required. 
- `make` is a built-in function to create a new map but has other uses too. 
