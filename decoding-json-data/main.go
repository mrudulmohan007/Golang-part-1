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

	DecodeJson()
}

//decoding of JSON

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJs Bootcamp",
                "Price": 4000,
                "website": "Youtube.com",
                "tags": ["React","JS"]
    }
	
	`)
	var allCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &allCourse)
		fmt.Printf("%#v\n", allCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	//some cases where u just want to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is: %T\n", k, v, v)
	}
}
