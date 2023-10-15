package merge_sort

import "testing"

func TestMergeSort(t *testing.T) {

	l := []int{55, 23, 43, 61, 12, 71, 95, 86, 31}
	e := []int{12, 23, 31, 43, 55, 61, 71, 86, 95}
	r := MergeSort(l)

	for i := range l {
		if e[i] != r[i] {
			t.Fail()
		}
	}
}
