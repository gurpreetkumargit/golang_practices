package main

import (
	"fmt"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {

	fmt.Println("channels in golang")

	myChan := make(chan int, 2)

	// fmt.Println(<-mychan)
	myChan <- 20
	wg.Add(2)
	go func(ch chan int) {
		fmt.Println("1", <-myChan)
		fmt.Println("2", <-myChan)

		wg.Done()
	}(myChan)

	go func(ch chan int) {

		myChan <- 10
		// myChan <- 11
		close(myChan)
		wg.Done()
	}(myChan)

	wg.Wait()
}
