//recover fn used to recover from panics
//only useful in deferred functions because when an app starts to panic
// it no longer execute the rest of the function but it will executes deferred functions
//so the proper place for using the recover fn is in deferred functions

package main

import (
	"fmt"
	"log"
)

// func main() {
// 	// fmt.Println("Start")
// 	// defer func() {
// 	// 	if err := recover(); err != nil { //recover function will return nil if the application
// 	// 		//is not panicking
// 	// 		log.Println("Error: ", err)
// 	// 	}
// 	// }()
// 	// panic("Something bad happened")
// 	// fmt.Println("End")

// 	// the output will be start and then Error: something bad happened
// 	//because here the recover return non-nill value means there is an error
// 	//so the error will be printed out

//

// }

//eg2

func main() {
	fmt.Println("Start")
	panicker()
	fmt.Println("End")
}

func panicker() {
	fmt.Println("About to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("Done panicking")

	//here firstly it will print start and then moves to panicker function
	//then it will print about to panic and program know its panicking
	// then recover function return a non-nil value so it will print Error:something bad happened
	//then panicker function stops and main function continues to execute and prints end
}
