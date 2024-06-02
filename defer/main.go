//defer is used to delay the execution of a statement untill fn exits
//defer statements doesn't execute at the end of the main function
//instead it executes after the main function but before the main function returned.
//Run in LIFO

package main

import "fmt"

func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle")
	fmt.Println("End") //output will be Start, End, Middle

	//in this eg if all these statements were defer go will follow LIFO(last in first out)
	// defer fmt.Println("Start")
	// defer fmt.Println("Middle")
	//defer fmt.Println("End")

	//the output will be: End,Middle,Start

	//another exmaple

	a := "start"
	defer fmt.Println(a)
	a = "end"

	//output will be start (because defer will take the argument when the defer is called)

}
