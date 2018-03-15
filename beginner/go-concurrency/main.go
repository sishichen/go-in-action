package main

import (
	"fmt"
	"sync"
)

func main() {
	// 0. data race
	// fmt.Println(getNumber())

	// 1. use waitGroup
	// fmt.Println(getNumberWithSync())

	// 2. use chan
	// fmt.Println(getNumberWithChan());

	// 3. use return chan
	i := <-getNumberWithReturnChan()
	fmt.Println(i)

	// 3.1 return chan with mutiple int
	// fmt.Println(getNumberWithReturnChan())
	for v := range getNumberWithReturnChan() {
		fmt.Println(v)
	}

	// --- with mutex
	// fmt.Println(getNumberWithMutex())
}

func getNumber() int {
	var i int
	go func() {
		i = 5
	}()

	return i
}

func getNumberWithSync() int {
	var i int
	// Initialize a waitgroup variable
	var wg sync.WaitGroup
	// `Add(1) signifies that there is 1 task that we need to wait for
	wg.Add(1) // you can try 2 and unmark second goroutine
	go func() {
		i = 5
		// Calling `wg.Done` indicates that we are done with the task we are waiting fo
		wg.Done()
	}()
	// go func() {
	// 	i = 7
	// 	// Calling `wg.Done` indicates that we are done with the task we are waiting fo
	// 	wg.Done()
	// }()

	// `wg.Wait` blocks until `wg.Done` is called the same number of times
	// as the amount of tasks we have (in this case, 1 time)
	wg.Wait()
	return i
}

func getNumberWithChan() int {
	var i int
	// Create a channel to push an empty struct to once we're done
	done := make(chan struct{})
	go func() {
		i = 5
		// Push an empty struct once we're done
		done <- struct{}{}
	}()
	// This statement blocks until something gets pushed into the `done` channel
	<-done
	return i
}

// return an integer channel instead of an integer
func getNumberWithReturnChan() <-chan int {
	// create the channel
	c := make(chan int)
	// 3.0
	go func() {
		// push the result into the channel
		c <- 5
	}()

	// 3.1
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		c <- i
	// 	}
	// 	close(c)
	// }()
	// immediately return the channel
	return c
}

// First, create a struct that contains the value we want to return
// along with a mutex instance
type SafeNumber struct {
	val int
	m   sync.Mutex
}

func (i *SafeNumber) Get() int {
	// The `Lock` method of the mutex blocks if it is already locked
	// if not, then it blocks other calls until the `Unlock` method is called
	i.m.Lock()
	// Defer `Unlock` until this method returns
	defer i.m.Unlock()
	// Return the value
	return i.val
}

func (i *SafeNumber) Set(val int) {
	// Similar to the `Get` method, except we Lock until we are done
	// writing to `i.val`
	i.m.Lock()
	defer i.m.Unlock()
	i.val = val
}

func getNumberWithMutex() int {
	// Create an instance of `SafeNumber`
	i := &SafeNumber{}
	// done:= make()
	// Use `Set` and `Get` instead of regular assignments and reads
	// We can now be sure that we can read only if the write has completed, or vice versa
	go func() {
		i.Set(3)
	}()

	// i.Set(3)
	// go func() {
	// 	i.Set(5)
	// }()

	return i.Get()
}
