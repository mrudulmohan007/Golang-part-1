package main

import (
	"fmt"
)

func main() {

	//syntax

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	//eg2

	// for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
	// 	fmt.Println(i, j)
	// }

	//eg3

	// for i := 0; i < 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	//eg4-looping over slices

	// s := []int{1, 2, 3, 4}
	// for k, v := range s {
	// 	fmt.Println(k, v)
	// }

	//eg5- Looping over maps

	// statePopulations := map[string]int{
	// 	"Adoor":        12345,
	// 	"Kottarakkara": 456789,
	// 	"Kottayam":     3456789,
	// 	"Kannur":       987654,
	// 	"Kollam":       5656777,
	// }
	// for k, v := range statePopulations {
	// 	fmt.Println(k, v)
	// }

	//if we need only value do like this:
	// for _,v:=range statePopulations{
	// 	fmt.Println(v)
	// }

	//eg-6- Looping over a string

	s := "hello!"
	for k, v := range s {
		fmt.Println(k, string(v))
	}

}
