package main

import "fmt"

// eg-1------------------------------------------------

type Engine interface {
	Start()
}

type PetrolEngine struct {
	Name string
}

func (e PetrolEngine) Start() {
	fmt.Println("starting...", e.Name)
}

type GasEngine struct {
	Name string
}

func (e GasEngine) Start() {
	fmt.Println("starting...", e.Name)
}

func Start(e Engine) {
	e.Start()
}
func main() {
	fmt.Println("Interfaces")
	pEngine := PetrolEngine{
		Name: "Petrol",
	}
	Start(pEngine)
	gEngine := PetrolEngine{
		Name: "Gas",
	}

	Start(gEngine)
}

// func main() {
// 	//eg-2
// 	fmt.Println("fmt") //interface don't describe data like structs but
// 	// it describe behaviors i.e,method definitions
// 	var runner Runner //object of Runner interface
// 	runner = Dog{Name: "Arjun"}
// 	runner.Run()
// }

// type Runner interface {
// 	Run() //behavior(method)

// }

// type Dog struct {
// 	Name string
// }

// func (d Dog) Run() { //implements behavior Runner
// 	fmt.Println(d.Name, "is running")
// }

// type Cat struct {
// 	Name string
// }

//eg3---------------------------------------------------
// Runner interface definition
// type Runner interface {
// 	Run()
// }

// // Dog struct definition
// type Dog struct {
// 	Name string
// }

// // Cat struct definition
// type Cat struct {
// 	Name string
// }

// // Run method implementation for Dog
// func (d Dog) Run() {
// 	fmt.Println(d.Name, "is running")
// }

// // Run method implementation for Cat
// func (c Cat) Run() {
// 	fmt.Println(c.Name, "is running")
// }

// // MakeRun function that accepts a Runner interface
// func MakeRun(runner Runner) {
// 	runner.Run()
// }

// // Main function
// func main() {
// 	var runner Runner           // Declare a variable of type Runner interface
// 	runner = Dog{Name: "Arjun"} // Assign a Dog instance to runner
// 	MakeRun(runner)             // Call MakeRun with the Dog instance

// 	var runner2 Runner         // Declare another variable of type Runner interface
// 	runner2 = Cat{Name: "car"} // Assign a Cat instance to runner2
// 	MakeRun(runner2)           // Call MakeRun with the Cat instance
// }
