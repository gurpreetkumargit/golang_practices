package main

import "fmt"

// collection or group of mix dataType
// overcome of the cons of map. in which we can only store one datatype of value

type address struct {
	name, city string
	pin_code   int
}

func main() {

	var a address
	fmt.Println(a)

	my_Address := address{"preet", "kaithal", 136027}
	fmt.Println(my_Address)

	address2 := address{name: "manik", city: "karnal", pin_code: 136065}
	fmt.Println("address 2: ", address2)

	//accessing the fields of struct
	fmt.Printf("name: %v, city: %v\n", my_Address.name, my_Address.city)

}
