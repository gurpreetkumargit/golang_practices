package main

import (
	"fmt"
	"time"
)

func show(str string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Second)
		fmt.Println(str)
	}
}
func main() {
	go show("lets go")

	show("hey")
}
