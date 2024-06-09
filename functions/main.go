package main

import "fmt"

// func main() {
// 	sayMessage("Hai my name is Mrudul Mohan")
// }
// //eg-1
// func sayMessage(msg string) {
// 	fmt.Println(msg)
// }

// eg-2
// func main() {
// 	for i := 0; i < 5; i++ {
// 		sayMessage("Hello Go!", i)
// 	}

// 	sayGreeting("Hello", "Mrudul")
// }
// func sayMessage(msg string, idx int) {
// 	fmt.Println(msg)
// 	fmt.Println("The value of the index is ", idx)
// }

// func sayGreeting(greeting, name string) { //here greeting variable is also of the type string
// 	fmt.Println(greeting, name)
// }

//eg-3
//passing in a pointer
// func main() {
// 	greeting := "Hello"
// 	name := "Stacey"
// 	sayGreeting(&greeting, &name)
// 	fmt.Println(name)
// }
// func sayGreeting(greeting *string, name *string) {
// 	fmt.Println(*greeting, *name)
// 	*name = "Ted"
// 	fmt.Println(*name)
// }

//eg-4-variadic parameter ----->> Variadic Function: A function that can take a variable
//number of arguments (values ...int)

// func main() {
// 	sum("The sum is ", 1, 2, 3, 4, 5)
// }
// func sum(msg string, values ...int) {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result = result + v
// 	}
// 	fmt.Println(msg, result)
// }

//eg-5- fn with arguments

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is ", s)
// }
// func sum(values ...int) int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result = result + v
// 	}
// 	return result
// }

//OR--ANOTHER WAY TO WRITE Eg-5---------------------------------------------------

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is ", s)
// }
// func sum(values ...int) (result int) {
// 	fmt.Println(values)
// 	for _, v := range values {
// 		result = result + v
// 	}
// 	return
// }

//eg-6- fn which return local variable as a pointer-------------------------------------------------

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is ", *s)
// }
// func sum(values ...int) *int {   //The function sum returns a pointer to an integer (*int).
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result = result + v
// 	}
// 	return &result
// }

//Eg-7 multiple return values

func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}
func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
