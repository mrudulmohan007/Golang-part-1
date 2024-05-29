package main

import "fmt"

func main() {
	fmt.Println("Switch examples are:--------------------------------------------------------->")
	i := 5
	//1.switch statements
	//condition based switch
	switch i {
	case 3:
		fmt.Println("Value is 3")
	case 5:
		fmt.Println("Value is 5")
	default:
		fmt.Println("Default value")
	}
	//logical switch
	switch {
	case i < 6:
		fmt.Println("i<6")
	case i > 6:
		fmt.Println("i>6")
	default:
		fmt.Println("Default value")
	}

	//3.goto statement
	c := 8
	if c > 7 {
		fmt.Println("c is > 7")
		goto mylabel
	}

mylabel:
	if c > 7 {
		fmt.Println("after label : c>7")
	}

	//another switch eg

	var k interface{} = 1.5
	switch k.(type) {
	case int:
		fmt.Println(" k is an int")
	case float64:
		fmt.Println("k is float64")
	case string:
		fmt.Println("k is string")
	default:
		fmt.Println("i is another type")

	}
}
