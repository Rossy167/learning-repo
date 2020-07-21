package main

import (
	"fmt"
	"sync"
	"time"
	// "runtime"
)

func main() {
	// uses goroutines rather than OS threads, are cheaper than OS threads and simpler to use, handled by go runtime
	// when goroutine waiting then the runtime swaps out for another routine in the same thread
	// obviously resources for that, but beats creating a new thread

	//can make em run parallel
	//runtime.GOMAXPROCS(2)

	var wait sync.WaitGroup
	wait.Add(2)

	go func() {
		defer wait.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("World")
	}()

	go func() {
		defer wait.Done()
		fmt.Println("Hello")
	}()

	wait.Wait()

	// try and find a way to use this in one of my portfolio projects, maybe messenger-cli

}
