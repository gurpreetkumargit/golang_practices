package main

import "fmt"

func main() {

	// an empty map
	// syntax
	// map[key_Type]value_Type{}

	// map with key and values
	// syntax
	// map[key_Type]value_Type{key1: value1, key2: value2}

	map_2 := map[int]string{

		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println(map_2)

	// using make method

	my_map := make(map[string]string)
	my_map["fistName"] = "preet"
	my_map["lastName"] = "dhiman"
	my_map["email"] = "preet@gmail.com"

	fmt.Println(my_map)

	// a list of map type

	listMap := make([]map[string]string, 0)

	userData := make(map[string]string)
	userData["name"] = "preet"
	userData["lastName"] = "dhiman"

	listMap = append(listMap, userData)
	fmt.Println(listMap)
}
