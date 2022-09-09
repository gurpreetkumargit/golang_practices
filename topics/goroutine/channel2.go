package main

import "fmt"

func ch(in chan int) {
	fmt.Println(<-in + 10)
}

func main() {

	fmt.Println("start")

	mych := make(chan int)
	go ch(mych)
	mych <- 30

	fmt.Println("end")
	// fmt.Println(<-mych)
}
