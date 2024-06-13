package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	go sayHello() //This line launches a new goroutine (lightweight thread) that
// 	//executes the sayHello function concurrently with the main function.
// 	time.Sleep(100 * time.Millisecond) //time.Sleep(100 * time.Millisecond) pauses the execution of the main function for 100 milliseconds.
// 	// This gives the sayHello goroutine time to run and print "Hello" before the main function completes and the program exits.
// }

// func sayHello() {
// 	fmt.Println("Hello")
// }

// NOTE:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// Flow of Execution---------------
// 1.The main function starts execution.
// 2.The go sayHello() statement launches the sayHello function in a new goroutine.
// 3.The main function then calls time.Sleep(100 * time.Millisecond), pausing its execution for 100 milliseconds.
// 4.During this pause, the sayHello goroutine runs and prints "Hello".
// 5.After 100 milliseconds, the main function resumes and finishes, causing the program to exit.

// **Key Points
// Goroutine: go sayHello() starts a new concurrent execution path.
// time.Sleep: Delays the main function to allow the goroutine to complete its task before the program exits.
// Concurrency: This example demonstrates basic concurrency in Go using goroutines.
// The time.Sleep is crucial here because, without it, the main function might finish and exit the program before the sayHello goroutine gets a chance to run.

//eg-2-The above program can be executed without the time.Sleep-----------------------
//thats y we use wait groups(wg)

// taskA is a function that prints messages and signals that it is done by calling wg.Done().
func taskA(wg *sync.WaitGroup) {
	fmt.Println("Task A")
	fmt.Println("Task A done")
	wg.Done() // Decrements the WaitGroup counter by one.
}

func main() {
	fmt.Println("concurrency")
	var wg sync.WaitGroup // Declares a WaitGroup.
	wg.Add(1)             // Increments the WaitGroup counter by one.
	go taskA(&wg)         // Starts taskA in a new goroutine, passing the WaitGroup pointer.
	wg.Wait()             // Blocks until the WaitGroup counter is zero.
	fmt.Println("Main Exit")
}

//Flow of Execution:

//** Main Function:-----------
// 1.Prints "concurrency".
// 2.Declares a WaitGroup named wg.
// 3.Calls wg.Add(1) to increment the counter by 1, indicating one goroutine is expected to complete.
// 4.Starts taskA as a new goroutine using go taskA(&wg).
// 5.Calls wg.Wait(), which blocks the main goroutine until the counter is decremented to zero.
// 6.After taskA finishes and calls wg.Done(), the counter decrements to zero, wg.Wait() unblocks, and "Main Exit" is printed.
