package maxmin

import "testing"

func TestMax(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, -1, 2, 4, 0, -3, 7, -3}
	expect := 7
	result := max(input...)
	if expect != result {
		t.Error("Wrong max value")
	}
}

func TestMin(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, -1, 2, 4, 0, -3, 7, -3}
	expect := -3
	result := min(input...)
	if expect != result {
		t.Error("Wrong min value")
	}
}

func TestMinNoArgs(t *testing.T) {
	if min() != 0 {
		t.FailNow()
	}
}

func TestMaxNoArgs(t *testing.T) {
	if max() != 0 {
		t.FailNow()
	}
}
