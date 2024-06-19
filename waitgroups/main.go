package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	websitelist := []string{

		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://youtube.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}
