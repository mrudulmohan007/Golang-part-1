package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Post req in golang")
	PerformPostJsonRequest()

}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:5500/post"

	//fake json  payload
	requestBody := strings.NewReader(`
{
	"Coursename":"Golang course",
	"price": 0,
	"platform":"Youtube",
}
	  
	 `)

	response, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
