package main

import "fmt"

type employee struct {
	name, city string
	age, id    uint
}

// method with receiver and employee type
func (a employee) show() {
	fmt.Println("name: ", a.name)
	fmt.Println("city: ", a.city)
	fmt.Println("age: ", a.age)
	fmt.Println("id: ", a.id)
}

func main() {
	// method is similar to func but diff is - method contains a receiver argument.
	// using this receiver argument we can access the prop of receiver
	// receiver can be a struct or non-struct type but
	// the receive must be present in the same package

	data := employee{
		name: "preet",
		city: "kaithal",
		age:  21,
		id:   1111,
	}

	data.show()
	fmt.Println(data)
}
