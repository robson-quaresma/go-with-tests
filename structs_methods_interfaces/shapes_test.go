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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got '%g' want '%g'", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectagle := Rectagle{12.0, 6.0}
		checkArea(t, rectagle, 72.0)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func TestAreaTDT(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectagle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got '%g' want '%g'", got, tt.want)
		}
	}
}
