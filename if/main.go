package main

import (
	"fmt"
)

func main() {
	fmt.Println("If in go------------------------------------")
	if true {
		fmt.Println("This test is true")
	}

	//eg1

	number := 50
	guess := 30
	if guess < number {
		fmt.Println("Too low")
	}
	if guess > number {
		fmt.Println("Too high")
	}
	if guess == number {
		fmt.Println("You got it !")
	}

	//eg2
	i := 5
	if i < 6 {
		fmt.Println("i<6")
	} else if i > 6 {
		fmt.Println("i>6")
	} else {
		fmt.Println("Failed")
	}

}
