// panic occur when program cannot continue at all
// Deferred functions will still execute
package main

import "fmt"

func main() {
	a, b := 1, 0
	ans := a / b
	fmt.Println(ans)

	//Output will be panic: runtime error

	//eg1

	fmt.Println("Start")
	panic("Something bad happened")
	fmt.Println("end") //output will be Start and a error saying: Something bad happened

	//eg2
	fmt.Println("Start")
	defer fmt.Println("This was deferred")
	panic("Something bad happened")
	fmt.Println("end") //Here output will be start and this was referred and panic: something bad happened
	//its because panic statements execute after deferred statements
	//so the order of execution is main function,deferred statements and then panic
	//and then finally return value

}
