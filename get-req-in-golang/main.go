package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("get req in golang")
	PerformGetRequest()

}
func PerformGetRequest() {
	const myurl = "http://webcode.me"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code is : ", response.StatusCode)
	fmt.Println("Content length is : ", response.ContentLength)
	content, _ := io.ReadAll(response.Body)
	fmt.Println("Content of the response is : ", string(content))

}
