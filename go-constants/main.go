package main

import "fmt"

func main() {
	const myConst int = 42
	fmt.Printf("Value of constant is %v and The datatype is %T\n", myConst, myConst)

	const a int = 14
	const b string = "football"
	const c float32 = 5.5
	const d bool = true
	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)

}
