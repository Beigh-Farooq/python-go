package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup //pointers

func main() {
	websitelist := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1) //This means that their is only one gorotine which you are going to fireup
	}
	wg.Wait() //This is responsible for Not allowing my main method to finish

}

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	} else {
		fmt.Printf("%d status code fo %s\n", res.StatusCode, endpoint)
	}
}
