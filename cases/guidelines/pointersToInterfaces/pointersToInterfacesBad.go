package pointersToInterfaces

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	rectangle := Rectangle{Width: 5, Height: 3}

	var shapePtr *Shape
	shapePtr = new(Shape)

	*shapePtr = &rectangle

	area := (*shapePtr).Area()

	fmt.Println("Area of the rectangle:", area)
}
