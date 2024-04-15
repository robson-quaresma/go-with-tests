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
	rectagle := Rectagle{12.0, 6.0}
	got := Area(rectagle)
	want := 72.0

	if got != want {
		t.Errorf("got '%.2f' want '%.2f'", got, want)
	}
}
