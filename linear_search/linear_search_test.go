package linear_search

import "testing"

func TestLinearSearch(t *testing.T) {

	v := 5
	l := []int{1, 3, 5, 9, 7, 0}
	e := 2
	r, err := LinearSearch(v, l)

	if err != nil && r != e {
		t.Fail()
	}
}
