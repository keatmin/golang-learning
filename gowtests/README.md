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
2. Pointers can be `nil`, so it's important to check if they are `nil` when a function/method returns one
3. Important to use pointers to mutate state or not make too many copies of large data
4. [ Handle error gracefully ](https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully)

## Maps
1. Use `err.Error()` to get string of err
2. t.Fatal() will exit and stop, useful to avoid panic in the next steps
3. An interesting property of maps is that you can modify them without passing as an address to it (e.g `&myMap`)
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
6. `delete` built-in function takes two args, map, and a key.

## Dependency Injection
1. Both bytes.Buffer and os.Stdout implements io.Writer
2. Understanding general purposes interface can make software reusable in a number of contexts
3. Essentially separating concern of the construction of an object

### WHY DI
Motivated by our tests we refactored the code so we could control where the data was written by injecting a dependency which allowed us to:

- **Test our code** If you can't test a function easily, it's usually because of dependencies hard-wired into a function or global state. If you have a global database connection pool for instance that is used by some kind of service layer, it is likely going to be difficult to test and they will be slow to run. DI will motivate you to inject in a database dependency (via an interface) which you can then mock out with something you can control in your tests.
- **Separate our concerns**, decoupling where the data goes from how to generate it. If you ever feel like a method/function has too many responsibilities (generating data and writing to a db? handling HTTP requests and doing domain level logic?) DI is probably going to be the tool you need.
- **Allow our code to be re-used in different contexts** The first "new" context our code can be used in is inside tests. But further on if someone wants to try something new with your function they can inject their own dependencies.

## Mocking
1. backtick syntax allows us to create string but includes newline
2. `i--` to loop backwards for `post statement` in `for statement`
3. when time.Sleep is added test will take 3s to run, quick feedback loop is needed and slow test ruins productivity
4. Spies are a kind of mock which can record how a dependency is used. They can record the arguments sent in, how many times it has been called, etc.

5. If refactoring causes a lot of tests to be changed, it's a sign of testing too much implementation details. Try to test for behaviour.

6. When writing a test, consider whether the test tests for behaviour or the implementation details? if the code needs to be refactored, does the test have to change?
```text
Although Go lets you test private functions, I would avoid it as private functions are implementation detail to support public behaviour. Test the public behaviour. Sandi Metz describes private functions as being "less stable" and you don't want to couple your tests to them.
```
7. When faced with less trivial examples, break the problem down into "thin vertical slices". Try to get to a point where you have working software backed by tests as soon as you can, to avoid getting in rabbit holes and taking a "big bang" approach.

## Concurrency
1.`()` at the end of `go` statement is so that func is executed at the same time they are declared
2. Maps in go don't like more than 1 thing trying to write to them at once. `Concurrent map writes` is the fatal error. This is a `race condition`. When output of software depends on timing and sequence of events
