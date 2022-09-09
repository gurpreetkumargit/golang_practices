package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)

	// receiving go routine
	wg.Add(2)

	go func(ch <-chan int) {
		fmt.Println("receive")
		i := <-ch
		fmt.Println(i)
		// ch <- 30
		wg.Done()
	}(ch)

	// for j := 1; j <= 5; j++ {

	// wg.Add(1)

	// sending go routing

	go func(ch chan<- int) {
		fmt.Println("send")

		ch <- 20

		// fmt.Println(<-)
		// fmt.Println(<-ch)
		wg.Done()

	}(ch)

	// }

	wg.Wait()
}
