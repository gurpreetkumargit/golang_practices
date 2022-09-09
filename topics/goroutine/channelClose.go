package main

import (
	"fmt"
	"sync"
)

func res(in chan string) {
	for i := 0; i < 6; i++ {
		in <- "preet"
		fmt.Println(<-in)
	}
	close(in)
}

func main() {
	wg := &sync.WaitGroup{}

	// var mych chan string

	mych := make(chan string)

	wg.Add(1)
	go func() {
		res(mych)
	}()

	wg.Wait()

	// for {
	// 	res, ok := <-mych

	// 	if ok != true {
	// 		fmt.Println("false")
	// 		break
	// 	} else {
	// 		fmt.Printf("true %v\n", res)
	// 	}

	// }

}
