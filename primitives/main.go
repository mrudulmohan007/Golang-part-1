package main

import "fmt"

func main() {
	var n bool = false
	fmt.Printf("value of n is %v,%T\n", n, n)
	p := 1 == 1
	q := 1 == 2
	fmt.Printf("Value of p is %v\n", p)
	fmt.Printf("Value of q is %v\n", q)

	var a bool //default value of a variable in go is zero(here its false)
	fmt.Printf("Value of a is %v\n", a)
}
