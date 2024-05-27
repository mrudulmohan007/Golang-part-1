// //Struct examples---------------------------------------------------------------

// package main

// import "fmt"

// //eg1
// type Film struct {
// 	collection int
// 	actorName  string
// 	companions []string
// }

// func main() {
// 	fmt.Println("Struct examples are as follows-------------------------- ")
// 	avesham := Film{
// 		collection: 100,
// 		actorName:  "Fahadh Faasil",
// 		companions: []string{
// 			"Amban",
// 			"Hipster ",
// 			"Bibimon",
// 		},
// 	}
// 	fmt.Println(avesham)
// 	fmt.Println(avesham.actorName)
// 	fmt.Println(avesham.companions)
// 	fmt.Println(avesham.companions[1])

// }

// 2ND EXAMPLE- achieving inheritance in go(composition through embedding)
package main

import "fmt"

type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal   //embedding
	SpeedKPH float32
	CanFly   bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name)

	//another way

	k := Bird{
		Animal:   Animal{Name: "Cuckoo", Origin: "Australia"},
		SpeedKPH: 45,
		CanFly:   true,
	}
	fmt.Println(k)
	fmt.Println(k.SpeedKPH)
}
