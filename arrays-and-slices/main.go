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

	//same eg but used "&" (pointer)

	c := [3]int{1, 2, 3}
	d := &c
	d[1] = 5
	fmt.Println(c)
	fmt.Println(b)

	fmt.Println("-----------SLICE-------------------------------------------")

	// "-----------SLICE-------------------------------------------"

	p := []int{1, 2, 3}
	q := p
	q[1] = 6
	fmt.Println(p)
	fmt.Println(q)
	fmt.Printf("Length: %v\n", len(p))
	fmt.Printf("Capacity: %v\n", cap(p))

	fmt.Println("-----------SLICE-------------------------------------------")

	//slice syntax

	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	all_elements := x[:]                        //slice of all elements
	slice_from_four_to_end := x[3:]             //slice from 4th element to end
	slice_first_6_elements := x[:6]             //slice first 6 elements
	slice_fourth_fifth_sixth_elements := x[3:6] //slice 4th,5th,6th elements

	fmt.Println(all_elements)
	fmt.Println(slice_from_four_to_end)
	fmt.Println(slice_first_6_elements)
	fmt.Println(slice_fourth_fifth_sixth_elements)

	fmt.Println("-----------SLICE-------------------------------------------")

	//slice eg:
	u := []int{}
	fmt.Println(u)
	fmt.Printf("Length: %v\n", len(u))
	fmt.Printf("Capacity: %v\n", cap(u))
	u = append(u, 1)
	fmt.Println(u)
	fmt.Printf("Length: %v\n", len(u))
	fmt.Printf("Capacity: %v\n", cap(u))

	u = append(u, []int{2, 3, 4, 5}...) // here ... is added to spread the slice to individual arguments
	//otherwise it will give u error
	fmt.Println(u)
	fmt.Printf("Length: %v\n", len(u))
	fmt.Printf("Capacity: %v\n", cap(u))

	fmt.Println("-----------SLICE-------------------------------------------")

	//another slice eg: print all elements of a slice

	i := []int{1, 2, 3, 4, 5}
	j := i[:]
	fmt.Println(j)
	//remove last element from a slice
	k := i[:len(i)-1]
	fmt.Println(k)

	//to remove an element from middle

	m := []int{1, 2, 3, 4, 5}
	n := append(m[:2], m[3:]...)

	fmt.Println(n)

}
