# Error Handling

## How do other languages do?
#### Java
```java
try {
   // happy flow
} catch(Exception e) {
   // error handling
}
```
#### Swift
```swift
// Option 1
do {
  // happy flow
} catch (let error as NSError) {
  // error handling
}

// Option 2: assign to an optional variable
let response = try? ETPhoneHome("Halp!")

// Option 3: crash the app if an error is thrown
let response = try! ETPhoneHome("BOOM")
```

#### C
```c
#define CHECK_OR_RETURN(contract, error_status, log_level, message_format, ...) \
if (!(contract)) { \
    global_logger.message(log_level, message_format, ##__VA_ARGS__); \
    return error_status; \
}
```
## And Go?
The error handling in Go doesn't have the concept like exception and try-catch, instead, you should explicitly return an `error` and handle it.

#### Go
```go
err := doSomething()
if err != nil {
  // error handling
}
// happy flow
```
or with multiple reutrn values
```go
result, err := doSomething()
if err != nil {
  // error handling
}
// happy flow
```

## What's the "error"
The `error` type is an interface with only one method: `Error`
```go
type error interface {
    Error() string
}
```

You could make an error by [errors.New](https://golang.org/src/errors/errors.go?s=293:320#L1):
```go
func doSomething() error {
	if condition {
		return errors.New("error la")
	}
	// ...
}
```

Also, Go allow you to defind your own error type by implementing the `error` interface:
```go
type Error string

func (e Error) Error() string { return string(e) }
```

## Are there something like try-catch in Go?
As we said, Go doen't have the same concept, Go programmers should (if need) handle the errors explicitly. But, there is still an alternative approach -- **recover** the **panic**
https://golang.org/ref/spec#Handling_panics

Don't abuse it. We should only use it when the panic is unexpected in normal flow (e.g., out of memory).


## References
- [Errors](https://github.com/golang/go/wiki/Errors)
- [Go's Error Handling is Elegant](https://davidnix.io/post/error-handling-in-go/)
- [Effective error handling in Go. - Morsing's blog](https://morsmachine.dk/error-handling)
- [Error handling patterns in Go](https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/)
- [Catching panics in Golang](https://stackoverflow.com/questions/25025467/catching-panics-in-golang)
- [Handling errors in C with “check-log-return”: why not use a macro?](https://softwareengineering.stackexchange.com/questions/335851/handling-errors-in-c-with-check-log-return-why-not-use-a-macro)
