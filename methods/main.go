// package main

// import "fmt"

// func main() {
// 	g := greeter{
// 		greeting: "Hello",
// 		name:     "Go",
// 	}
// 	g.greet()
// }

// type greeter struct {
// 	greeting string
// 	name     string
// }

//method

// func (g greeter) greet() {
// 	fmt.Println(g.greeting, g.name)
// }

//eg-2-------------

// package main

// import "fmt"

// // Define a struct type called 'rectangle'
// type rectangle struct {
// 	width  int
// 	height int
// }

// // Method for the 'rectangle' struct to calculate the area
// // The receiver (r rectangle) allows the method to access fields of the struct
// func (r rectangle) area() int {
// 	return r.width * r.height
// }

// func main() {
// 	// Create an instance of the 'rectangle' struct
// 	r := rectangle{width: 10, height: 5}

// 	// Call the 'area' method on the 'rectangle' instance
// 	area := r.area()

// 	// Print the area
// 	fmt.Println("Area of the rectangle:", area)
// }

//EG-3----------------------------------------------------------------------------------------------------------------

// Key Differences:
// Association:

// Functions are independent and not tied to any type.
// Methods are tied to a receiver type, typically a struct, and can access its fields and methods.
// Syntax:

// Functions do not have a receiver parameter.
// Methods have a receiver parameter which is specified before the method name.
// Example to Illustrate Both:

package main

import "fmt"

// Function
func add(a int, b int) int {
	return a + b
}

// Struct
type Rectangle struct {
	width, height int
}

// Method associated with Rectangle struct
func (r Rectangle) area() int {
	return r.width * r.height
}

func main() {
	// Using the function
	sum := add(3, 4)
	fmt.Println("Sum:", sum)

	// Using the method
	rect := Rectangle{width: 10, height: 5}
	area := rect.area()
	fmt.Println("Area of the rectangle:", area)
}
