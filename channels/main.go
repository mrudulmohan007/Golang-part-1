package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Golang")
	myCh := make(chan int, 2) //buffer channel
	wg := &sync.WaitGroup{}
	// myCh <- 5 //push a value into myCh
	// fmt.Println(<-myCh)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)

		defer wg.Done()
	}(myCh, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		myCh <- 7
		close(myCh) //closing the channel
		defer wg.Done()
	}(myCh, wg)
	wg.Wait()
}
