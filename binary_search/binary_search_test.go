package binary_search

import "testing"

func TestBinarySearch(t *testing.T) {

	l := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 13}
	v := 11
	r := BinarySearch(v, l)

	if !r {
		t.Fail()
	}
}
