https://divan.github.io/posts/go_concurrency_visualize/
> Read the article and check the gif it provided

# Go Concurrency Visualize
* `GOMAXPROCS=4`: sets the maximum number of CPUs that can be executing simultaneously.
* Goroutines leak
* Parallelism is not Concurrency
  * Parallelism is simply running things in parallel.
  * Concurrency is a way to structure your program.
  * Concurrent program may or may not be parallel

# Bonus: Go Trace
* import `runtime/trace`
* Call `trace.Start(f)` and `trace.Stop` between the code you want to trace. e.g: 

```
func main() {
	f, err := os.Create("trace.out")
	defer f.Close()

	err = trace.Start(f)
	defer trace.Stop()

	fmt.Println("hello")
}
```

* `go tool trace trace.out`

Ref: 
* https://divan.github.io/posts/go_concurrency_visualize/
* https://making.pusher.com/go-tool-trace/