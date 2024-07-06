package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // `` is used to give alias to fields in struct
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // here ``encrypts password
	Tags     []string `json:"tags,omitempty"` //use omitempty if any field is nil it wont show up
}

func main() {
	fmt.Println("Welcome to JSON tutorial")
	EncodeJson()
}

//encoding of JSON

func EncodeJson() {
	ytCourses := []course{
		{
			"ReactJs Bootcamp", 4000, "Youtube.com", "qwerty", []string{"React", "JS"},
		},
		{
			"Golang Bootcamp", 5000, "Udemy.in", "asdf", []string{"Go", "Backend"},
		},
		{
			"AWS Bootcamp", 2000, "Youtube", "dfgh", nil,
		},
	}

	//package this data to JSON data
	finalJson, err := json.MarshalIndent(ytCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}
