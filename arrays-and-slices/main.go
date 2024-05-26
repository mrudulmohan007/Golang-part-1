package main

import "fmt"

func main() {
	fmt.Println("-----------ARRAYS-------------------------------------------")

	grades := [3]int{97, 585, 93}
	fmt.Printf("Grades: %v", grades)

	fmt.Println("----------------------------------------------")

	var students [5]string
	fmt.Printf("Students: %v\n", students)
	students[0] = "Mrudul"
	students[1] = "Vinay"
	students[2] = "Abel"
	students[3] = "yasin"
	students[4] = "melvin"
	fmt.Printf("Students: %v\n", students)
	fmt.Printf("Student number 1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))

	fmt.Println("----------------------------------------------")

	//array of array

	var identityMatrix [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{1, 0, 0}, [3]int{1, 0, 0}}
	fmt.Println(identityMatrix)

	//or

	var identityMatrix2 [3][3]int
	identityMatrix2[0] = [3]int{1, 0, 0}
	identityMatrix2[1] = [3]int{0, 1, 0}
	identityMatrix2[2] = [3]int{0, 0, 1}

	fmt.Println(identityMatrix2)

	//array eg:

	a := [3]int{1, 2, 3}
	b := a
	b[1] = 5
	fmt.Println(a)
	fmt.Println(b)

	//same eg but used &

	c := [3]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(b)

}
