package treesort

import "testing"

func TestSort(t *testing.T) {
	arr := []int{10, 15, 7, 78, 11}
	result := []int{7, 10, 11, 15, 78}
	Sort(arr)
	if !compareSlices(arr, result) {
		t.Fail()
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
