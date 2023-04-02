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
