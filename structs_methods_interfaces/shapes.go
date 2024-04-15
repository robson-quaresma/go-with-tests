package structsmethodsinterfaces

import "math"

type Rectagle struct {
	Width  float64
	Height float64
}

func (r Rectagle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func Perimeter(rectangle Rectagle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectagle) float64 {
	return rectangle.Width * rectangle.Height
}
