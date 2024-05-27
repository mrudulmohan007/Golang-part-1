package main

import "fmt"

func main() {
	//maps examples
	fmt.Println("Maps examples are as follows ")

	statePopulations := map[string]int{
		"California":     1000000,
		"Texas":          1234556,
		"Florida":        34567,
		"New York":       5678956,
		"Adoor":          12345,
		"pathanamthitta": 345678,
		"kollam":         8765678,
	}
	fmt.Println(statePopulations)
	fmt.Println(statePopulations["Adoor"])
	statePopulations["Georgia"] = 1098543 //adding a key to the map
	fmt.Println(statePopulations)
	delete(statePopulations, "Georgia") // delete a key from the map

	pop, ok := statePopulations["Ohio"] //check whether a key is present in the map
	fmt.Println(pop, ok)

	//length of the map

	fmt.Println(len(statePopulations))

}
