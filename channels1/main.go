package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Channels tutorial-2")
	ch := make(chan int)
	wg.Add(2)
	go func() {
		i := <-ch //receiving data from the channel
		defer wg.Done()
		fmt.Println(i)

	}()
	go func() {
		defer wg.Done()
		ch <- 42 //assigning value to the channel
	}()
	wg.Wait()

}
