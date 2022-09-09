package main

import (
	"fmt"
	"math"
)

type circleInterface interface {
	circleArea() float64
	circleParam() float64
}

type circle struct {
	radius float64
}

func (c circle) circleArea() float64 {
	fmt.Println("area of circle: ")
	return math.Pi * c.radius * c.radius
}

func (cp circle) circleParam() float64 {
	return 2 * math.Pi * cp.radius
}

// type rectangle struct {
// 	length, width float64
// }

// func (r rectangle) area() float64 {
// 	fmt.Println("area of rectangle: ")
// 	return r.length + r.width
// }

// func measure(g geometry) {
// 	fmt.Println(g.area())
// 	fmt.Println(g.circleParam())
// }

func main() {
	// my_Rect := rectangle{10, 20}
	my_Circle := circle{12}

	fmt.Println("circle Data: ", my_Circle)
	// fmt.Println("rectangle data: ", my_Rect)

	fmt.Println("circle area: ", my_Circle.circleArea())
	// fmt.Println("rectangle area: ", my_Rect.area())
	fmt.Println("circle parameter: ", my_Circle.circleParam())

	fmt.Println("\n \n")

	var cr circleInterface

	cr = circle{12.7}
	fmt.Println(cr.circleArea())
	fmt.Println(cr.circleParam())
	fmt.Println(cr)

}
