package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

var signals = []string{"test"}

func main() {

	webSiteList := []string{
		"https://google.com",
		"https://preetdhiman.com",
		"https://amazon.com",
		"https://spreadtech.com",
	}

	for _, web := range webSiteList {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)
}

func getStatusCode(endpoint string) {

	res, err := http.Get(endpoint)

	defer wg.Done()

	if err != nil {
		fmt.Printf("OOPS in endpoint for %v\n", endpoint)
	} else {

		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %v\n", res.StatusCode, endpoint)
	}
}
