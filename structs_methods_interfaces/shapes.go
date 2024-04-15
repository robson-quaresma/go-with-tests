package structsmethodsinterfaces

type Rectagle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectagle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectagle) float64 {
	return rectangle.Width * rectangle.Height
}
