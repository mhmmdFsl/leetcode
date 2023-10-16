package pascal_triangel

import (
	"reflect"
	"testing"
)

func TestPascalTriangel(t *testing.T) {

	rn := 5
	pt := PascalTriangel(rn)
	var e [][]int
	e = append(e, []int{1})
	e = append(e, []int{1, 1})
	e = append(e, []int{1, 2, 1})
	e = append(e, []int{1, 3, 3, 1})
	e = append(e, []int{1, 4, 6, 4, 1})
	compareSlices(e, pt, t)
}

func compareSlices(slice1, slice2 [][]int, t *testing.T) {
	// Check if the dimensions match
	if len(slice1) != len(slice2) {
		t.Fail()
	}

	for i := 0; i < len(slice1); i++ {
		// Check if the lengths of the inner slices match
		if !reflect.DeepEqual(slice1[i], slice2[i]) {
			t.Fail()
		}
	}
}
