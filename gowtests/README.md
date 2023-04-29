# GO with tests
## Hello World
1. Write a test
2. Make the compiler pass
3. Run the test, see that it fails and check the error message is meaningful
4. Write enough code to make the test pass
5. Refactor

## Iter
1. `go test -bench=.` to run benchmarks

## Arr
1. Question value of tests because tests have cost.
2. `go test -cover` to see coverage
3. Making a copy of a slice from another slice helps with garbage collection of a bigger array


## Pointers
1. If you see a function that takes arguments or returns values that are interfaces, they can be nillable. `Error` is an interface
2. Pointers can be `nil` so it's important to check if they are `nil` when a function/method returns one
3. Important to use pointers to mutate state or not make too many copies of large data
4. [ Handle error gracefully ](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

## Maps
1. Use err.Error() to get string of err
2. t.Fatal() will exit and stop, useful to avoid panic in the next steps
3. An interesting property of maps is that you can modify them without passing as an address to it (e.g &myMap)
4. A `nil` map behaves like an empty `nil` but attempting to write to a `nil` map will cause a panic. Never initialize an empty map variable
```golang
var m map[string]string
```

initialize an empty map like below
```golang
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```
5. Make errors a `const`, and implements own `Err` type with `Error()` interface. [ This is why ](https://dave.cheney.net/2016/04/07/constant-errors)
