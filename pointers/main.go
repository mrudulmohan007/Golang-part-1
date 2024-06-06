package main

import "fmt"

func main() {
	//eg1

	// var a int = 42
	// var b *int = &a  //here b is pointing to a ,means b holds the memory location of a
	// fmt.Println(a, b)

	//output will be:
	// 42 0xc000016088 because b has the memory location of a
	//to print the value of that memory location which is taken by b:

	var a int = 42
	var b *int = &a
	fmt.Println(a, *b)
	//output will be 42 42
	//here *b means dereference means it asks go runtime to find the memory location of b
	//and pull the value from there
	//now if we change the value of a , the value of b will also be changed

	*b = 14
	fmt.Println(a, *b) //now value of both a and b will be changed

	//eg2- Pointer Arithmetic

	c := [3]int{1, 2, 3}
	d := &c[0]                        //d is pointing to c[0] means the memory address of 1
	e := &c[1]                        //e is pointing to c[1] means  memory address of 2
	fmt.Printf("%v %p %p\n", c, d, e) //this will print the array c, the memory address of d and e
	//to get the value of d and e  try this: fmt.Println(c, *d, *e)

	//eg3

	var ms *myStruct // This declares a variable ms of type *myStruct,
	// which means ms is a pointer to a myStruct.
	ms = &myStruct{foo: 42} //&- address of operator
	fmt.Println(ms)         //This prints the value of ms, which is the address of the
	//myStruct instance it points to.

	// 	ms is a pointer variable that holds the address of a myStruct instance.
	// An object of myStruct is created with foo set to 42, and ms points to this object.

}

type myStruct struct {
	foo int
}

//note: important!!!---------------------------------------------------------------------------------------
// a:= [3]int{1,2,3}
// b:=a
// fmt.println(a,b)
// a[1]=42
// fmt.println(a,b)

// value of a and b will be different here because when we assigned
//b:=a, the value of array a copied to b
// but if a was a slice, internal representation of a slice actually
//has a pointer to the array , when we do b:=a effect of copying is bit different
//i.e,slice is actually projection of an underlying array,
//so the slice doesn't contain the data itself, the slice contains
// a pointer to the array. So when we copy(b:=a),part of data that gets copied is
// a pointer. Maps also has this behavior but in primitives,arrays structs
//this wont happen
