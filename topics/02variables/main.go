package main

import "fmt"

const LoginToken = "wdkjsbslknc"

func main() {

	// string value

	var name string = "preet"
	fmt.Println(name)
	fmt.Printf("the dataType is: %T \n", name)
	fmt.Print("his name is: " + name + "\n")

	// boolean

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)

	//int   int dataType can store +ve and -ve both values

	var smallValue int = 256
	fmt.Println(smallValue)
	fmt.Printf("%T \n", smallValue)
	var nagValue int32 = -283490
	fmt.Println(nagValue)
	fmt.Printf("%T \n", nagValue)

	//uint  uint can store only 0 and +ve values

	var smallUint uint16 = 23123
	fmt.Println(smallUint)
	fmt.Printf("dataType is: %T \n", smallUint)

	//float values

	var floatVal float32 = 12.687218921
	var floatVal2 float64 = 12.68721892113
	fmt.Println(floatVal)
	fmt.Println(floatVal2)

	// some default values and aliases

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("dataType is %T \n", anotherVar)
	anotherVar = 3211
	fmt.Println(anotherVar)

	// implicit way to declare var

	var website = "google"
	fmt.Println(website)

	//no var style

	secondVar := 13423
	fmt.Println(secondVar)
	fmt.Printf("dataType is %T \n", secondVar)

	//const outside the main func() and Public const

	fmt.Println(LoginToken)
	fmt.Printf("dataType is LoginToken is: %T \n", LoginToken)
}
