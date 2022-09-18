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
- Printf() formats the string without adding new lines. The same pattern for functions when it's log.Printf() or fmt.Errorf(). Same pattern can be observed for Println(), where a new line is created. 
- `err` is usually returned in a function, if `err` is `nil` then the function returns successfully

Types of verbs in go that is used for conversions:
```
%d decimal integer
%x, %o, %b integer in hexadecimal, octal, binary
%f, %g, %e floating point number
%t boolean
%c rune
%s string
%q quoted string "abc"
%v any value in a natural format
%T type of any value
%% literal percent design
```

### Chapter 1.5
[Fetch](gopl/ch1/fetch.go)
- `net` package make it easy to send and receie information through the internet, make low-level network connections, setup servers. 
- `fetch` is used to fetch content of URL, inspired by `curl` 

- Function call `io.Copy(dst, src)` reads from src and writes to dst. Using this instead of ioutil.ReadAll to copy response body to os.Stdout without requiring a buffer large enough to hold the entire stream. Always check for err.
- Use `strings.HasPrefix` to see if a string has a specific prefix
- `resp.status` also prints the HTTP status code

### Chapter 1.6
[FetchAll](gopl/ch1/fetchall.go)
- `goroutine` is concurrent function execution. 
- `channel` is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine
- `go` statement creates additional goroutines. 

```
“When one goroutine attempts a send or receive on a channel, it blocks until another goroutine attempts the corresponding receive or send operation, at which point the value is transferred and both goroutines proceed”

Excerpt From: Brian W. Kernighan. “The Go Programming Language (Addison-Wesley Professional Computing Series)”
```
### Chapter 1.7
[Server1](gopl/ch1/server1.go)
[Server2](gopl/ch1/server2.go)
[Server3](gopl/ch1/server3.go)
- Using `sync.Mutex` to Lock and Unlock to ensure one goroutine can access variable one at a time 
- Go allows local variable declaration to precede if condition
```golang
err := r.ParseForm()
if err != nil {
  bla
}
```
can be instead be written as such reduces the scope of `err`
```golang
if err := r.ParseForm(); err != nil {
  bla
}
```

### Chapter 1.8
- Switch statement evaluate results from top to bottom
```golang
switch flipcoin() {
  case "heads":
    heads++
  case "tails":
    tails++
  default:
    fmt.Println("landed on edge")
  }
}
```
- default case matches if none of the other does; it may be placed anywhere.

`switch` can be used without an operand as such, a `tagless switch`
```golang
func bla(x int) int {
  switch {
    case x > 0: 
      return +1
    default:
      return 0
    case x < 0: 
      return -1
  }
}
```
- `pointers` “The & operator yields the address of a variable, and the * operator retrieves”

