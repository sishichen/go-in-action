# Before We Start...
1. `goroutine`: a light weight thread
  * `go f(x, y, z)`
2. `channel`: data structures that enable safe data communication between goroutines
```
channelForInts := make(chan int) // Create channel
channelForInts <- 5              // Put 5 into channel
number := <-channelForInts       // Get 5 from channel
```
  * `r := make(<-chan bool)`: can only read from
  * `w := make(chan<- []os.FileInfo)`: can only write to

# Let's Talk About Channel
* Metaphor: Think `channel` as a Portal[2]

## Synchronous Channel
* Unbuffered channel: `done := make(chan bool)`
* All operations on unbuffered channels block the execution until both sender and receiver are ready to communicate

```
func main() {
	done := make(chan bool)

	go func() {
		println("goroutine message")
		// Tell the main function everything is done.
		// This channel is visible inside this goroutine because
		// it is executed in the same address space.
		done <- true
	}()

	println("main function message")
	<-done
}
```

## Asynchronous Channel
* Channel with buffer

```
func main() {
	message := make(chan string) // no buffer, bloooooooocked
	// message := make(chan string, 3) // with buffer
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			fmt.Println("send message")
			message <- fmt.Sprintf("message %d", i)
		}
	}()

	time.Sleep(time.Second * 3)
	for i := 1; i <= count; i++ {
		fmt.Println(<-message)
	}
}
```

## Deadlock

```
func main() {
	c := make(chan int)
	c <- 42    // write to a channel, blocks the executions, QQ
	val := <-c // read from a channel
	println(val)
}
```

* the right way: 
```
func main() {
	c := make(chan int)
	go func() {
		c <- 42 // write to a channel
	}()
	val := <-c
	println(val)
}
```

* Or you can set the channel buffer, then it won't be deadlock

## Range Expression
* `for k := range myArray`, any iterateable object
* For channels, the iteration proceeds until channel closed

```
func main() {
	message := make(chan string)
	count := 3

	go func() {
		for i := 1; i <= count; i++ {
			message <- fmt.Sprintf("message %d", i)
		}
		// close(message) // close the channel
	}()

	for msg := range message { // work until the channel is closed
		fmt.Println(msg)
	}
}
```

* Read value on closed channels will return default value for the channel type
```
func main() {
	message := make(chan bool)
	close(message)
	println(<-message)	// return false
}
```

* e.g: no need to put value in channel "done"
  
```
func main() {
	done := make(chan bool)

	go func() {
		println("goroutine message")
		close(done)
	}()

	println("main function message")
	<-done
}
```

## Multiple Channels and Select
* Let's take a look for a complex example

```
func getMessagesChannel(msg string, delay time.Duration) <-chan string {
	c := make(chan string)
	go func() {
		for i := 1; i <= 3; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			// Wait before sending next message
			time.Sleep(time.Millisecond * delay)
		}
	}()
	return c
}

func main() {
	c1 := getMessagesChannel("first", 300) // bottleneck
	c2 := getMessagesChannel("second", 150)
	c3 := getMessagesChannel("third", 10)

	for i := 1; i <= 3; i++ {
		println(<-c1)
		println(<-c2)
		println(<-c3)
	}
}
```

* Use `select`
```
func main() {
	c1 := getMessagesChannel("first", 300) // bottleneck
	c2 := getMessagesChannel("second", 150)
	c3 := getMessagesChannel("third", 10)

	for i := 1; i <= 9; i++ {
		select {
		case msg := <-c1:
			println(msg)
		case msg := <-c2:
			println(msg)
		case msg := <-c3:
			println(msg)
		}
	}
}
```

# Share Memory By Communicating
* Traditional treading models: communicate between threads using shared memory
  * Shared data structures are protected by locks
  * Thread-safe data structures
* Go: Communicating Sequential Processes
  * Channel pass data between goroutines
> Do not communicate by sharing memory; instead, share memory by communicating.

# Upcoming
* Effective go: https://golang.org/doc/effective_go.html

Ref: 
1. http://guzalexander.com/2013/12/06/golang-channels-tutorial.html
2. http://blog.mergermarket.it/now-youre-thinking-with-channels/
3. https://blog.golang.org/share-memory-by-communicating