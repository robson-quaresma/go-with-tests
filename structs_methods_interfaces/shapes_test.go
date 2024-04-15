package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectagle := Rectagle{10.0, 10.0}
	got := Perimeter(rectagle)
	want := 40.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectagle := Rectagle{12.0, 6.0}
		got := rectagle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got '%g' want '%g'", got, want)
		}
	})

}
