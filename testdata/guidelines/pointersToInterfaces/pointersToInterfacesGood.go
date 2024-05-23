package pointersToInterfaces

import "fmt"

type Shape2 interface {
	Area() float64
}

type Rectangle2 struct {
	Width  float64
	Height float64
}

func (r Rectangle2) Area() float64 {
	return r.Width * r.Height
}

func mainPointersGood() {
	rectangle := Rectangle2{Width: 5, Height: 3}

	var shapePtr Shape2

	shapePtr = rectangle

	area := shapePtr.Area()

	fmt.Println("Area of the rectangle:", area)
}
