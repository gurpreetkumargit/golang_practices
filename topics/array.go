package main

import "fmt"

func main() {

	// var new_array = [5]string{}
	// new_array[0] = "preet"
	// new_array[1] = "dhiman"
	// new_array[2] = "yes"
	// new_array[3] = "ok"
	// new_array[4] = "done"

	new_array := [5]string{"preet", "dhiman", "yes", "ok", "done"}

	fmt.Println(new_array)
	fmt.Println(new_array[2])
	fmt.Println(len(new_array))
	fmt.Println(cap(new_array))

	//change element of array

	new_array[0] = "gurpreet"

	fmt.Println(new_array[0])

	//initialize only specific element
	newArray := [6]int{1: 32, 5: 43}

	fmt.Println(newArray)
	fmt.Println(len(newArray))
	fmt.Println(cap(newArray))
}
