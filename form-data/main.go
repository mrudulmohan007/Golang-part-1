package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println("Form data in go")
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myurl = "https://httpbin.org/post"

	//form data
	data := url.Values{}
	data.Add("Firstname", "Mrudul")
	data.Add("Laststname", "Mohan")
	data.Add("Email", "mrudul@gmail.com")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}
