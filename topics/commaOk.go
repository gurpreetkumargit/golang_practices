package main

import "fmt"

func main() {

	myMap := map[string]string{
		"name":     "preet",
		"lastName": "dhiman",
		"city":     "kaithal",
	}

	// accessing data from map
	var fName = myMap["name"]
	var lName = myMap["invalid"]

	fmt.Printf("first name: %v\nlast Name: %v\n", fName, lName)

	// multiple assignment
	// if the index value found in map is valid then ok will true otherwise false
	var fName1, ok1 = myMap["name"]
	var city, ok2 = myMap["invalid"]

	fmt.Printf("value of fName1: %v | value of ok1: %v\n", fName1, ok1)
	fmt.Printf("value of city: %v | value of ok2: %v\n", city, ok2)

	// // comma ok idiom

	if fName2, ok3 := myMap["name"]; ok3 {
		fmt.Println("comma ok true")
		fmt.Println(fName2)
	}
}
