package structsmethodsinterfaces

type Rectagle struct {
	Width  float64
	Height float64
}

func (r Rectagle) Area() float64 {
	return 0
}

type Circle struct {
	Radius float64
}

func (r Circle) Area() float32 {
	return 0
}

func Perimeter(rectangle Rectagle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectagle) float64 {
	return rectangle.Width * rectangle.Height
}
