package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = &sync.WaitGroup{}

func main() {

	ch := make(chan int, 5)

	wg.Add(2)
	go func(ch <-chan int) {
		fmt.Println("receive")

		t := time.Now()
		fmt.Println(t)
		// i := <-ch
		// fmt.Println(i)
		// i = <-ch
		// fmt.Println(i)
		// fmt.Println(<-ch)

		fmt.Println("using range loop")
		// time.Sleep(2 * time.Second)
		for i := range ch {
			fmt.Println(i)
		}

		// instead of using for range loop we will use ok syntax to check the sending go is close or not

		// for {
		// 	if i, ok := <-ch; ok {
		// 		fmt.Println(i)
		// 	} else {
		// 		fmt.Println("send channel closed")
		// 		break
		// 	}
		// }
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		fmt.Println("send")
		ch <- 12
		ch <- 10
		// fmt.Println(<-ch)
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}
