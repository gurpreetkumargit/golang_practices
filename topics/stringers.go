package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

// type person struct {
// 	name   string
// 	age    int
// 	height float32
// }

// func (p person) String() string {
// 	return fmt.Sprintf("name: %v, age : %v and (height: %v) \n", p.name, p.age, p.height)
// }

func (op IPAddr) string() string {
	opStr := []string{}

	for _, val := range op {
		opStr = append(opStr, fmt.Sprint(int(val)))
	}
	return strings.Join(opStr, ".")
}

func main() {

	// var a = person{
	// 	"preet", 21, 5.5,
	// }

	// fmt.Println(a)

	mp := map[string]IPAddr{
		"ids":  {1, 2, 3, 5},
		"days": {2, 1, 5, 3},
	}

	for name, ip := range mp {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
