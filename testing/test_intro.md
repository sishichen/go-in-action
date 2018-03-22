## Test intro
#### An Introduction to Testing in Go | TutorialEdge.net
https://tutorialedge.net/golang/intro-testing-in-go/

### An Example of how your project would be structured
```
myproject/
- calc.go
- calc_test.go
- main.go
- main_test.go
```

### run test
```
$ go test 
```

### run verbose test
```
$ go test -v
```

### note
1. test file can access function at files in same folder location 