# Error Handling

## Are there something like try-catch in Go?
There is a similar approach -- **recover**
https://golang.org/ref/spec#Handling_panics

Don't abuse it. We should only use it when the panic is unexpected in normal flow (e.g., out of memory).


## References
- [Errors](https://github.com/golang/go/wiki/Errors)
- [Go's Error Handling is Elegant](https://davidnix.io/post/error-handling-in-go/)
- [Effective error handling in Go. - Morsing's blog](https://morsmachine.dk/error-handling)
- [Error handling patterns in Go](https://mijailovic.net/2017/05/09/error-handling-patterns-in-go/)
- [Handle HTTP Request Errors in Go @ Alex Pliautau's Blog](http://pliutau.com/handle-http-request-errors-in-go/)
- [Catching panics in Golang](https://stackoverflow.com/questions/25025467/catching-panics-in-golang)
