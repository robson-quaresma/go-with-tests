package arraysslices

import "testing"

func TestSum(t *testing.T) {
	// explict array declaration
	// numbers := [5]int{1, 2, 3, 4, 5}

	// implicit array declaration
	numbers := [...]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
	}
}
