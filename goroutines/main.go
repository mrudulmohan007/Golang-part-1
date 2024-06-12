package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello() //This line launches a new goroutine (lightweight thread) that
	//executes the sayHello function concurrently with the main function.
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}
