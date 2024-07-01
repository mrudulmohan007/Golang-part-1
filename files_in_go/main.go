package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files in Go")
	content := "This is a test file that should be put in a text file"
	//we have to put this content in a text file
	file, err := os.Create("./test.txt")
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./test.txt")
}

// fn to read the file
func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}

//fn to check error

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
