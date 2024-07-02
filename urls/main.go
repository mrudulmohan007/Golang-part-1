package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.flipkart.com/viewcart?exploreMode=true&preference=FLIPKART"

func main() {
	fmt.Println("Welcome to handling url in golang")

	//parsing the url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)
	fmt.Println(qparams["preference"])

	for _, value := range qparams {
		fmt.Println("Param is:", value)
	}

	//constructing a url

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host:   "www.flipkart.com",
		Path:   "/viewcart",
	}
	fmt.Println(partsOfUrl)

}
