package main

import "fmt"

var a int = 40
var (
	name     string = "Mrudul"
	lastName string = "Mohan"
	place    string = "Pathanamthitta"
)

func main() {

	fmt.Printf("My name is %v, and my place is %v\n", name, place)

	//type1-used when we have to assign the variable in a loop or something
	var i int
	i = 10
	fmt.Println("Value of i is :", i)
	//type2
	var j string = "Mrudul"
	fmt.Println("Value of j is :", j)
	//type3
	k := 7
	fmt.Println("Value of k is :", k)

	fmt.Printf("value of a is : %v", a)

}
