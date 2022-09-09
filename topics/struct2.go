package main

import "fmt"

//normal struct

type userData struct {
	firstName, lastName, city string
	age                       uint
}

type db struct {
	id  uint
	db1 userData
}

// we will learn here nested struct , anonymous struct and field

func main() {

	// accessing the fields of struct

	user1 := db{
		id:  101,
		db1: userData{"preet", "dhiman", "kaithal", 21},
	}
	fmt.Println(user1)

	// accessing fields from a nested struct
	fmt.Println(user1.db1.age)

	// anonymous structure
	fmt.Println("anon struct")
	anonStruct := struct {
		name, city string
	}{
		"preet", "kaithal",
	}

	fmt.Println(anonStruct)

	// access fields of anon struct
	fmt.Println(anonStruct.name)

	// anon fields in struct

	type anonFields struct {
		int
		string
	}

	anonData := anonFields{
		12, "ok",
	}

	fmt.Println("anonFields: ")
	fmt.Println(anonData)
	fmt.Println(anonData.int)
	fmt.Println(anonData.string)

}
