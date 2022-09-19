package main

import "fmt"

func index[T comparable](s []T, x T) int {
	fmt.Println("x is: ", x)

	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func main() {

	in := []int{10, 11, 87, 18, 6}

	op := index(in, 87)

	fmt.Println("match index is: ", op)
}
