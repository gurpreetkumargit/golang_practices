package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome")

	go func() {
		fmt.Println("this is anonymous goroutine")
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("end")
}
