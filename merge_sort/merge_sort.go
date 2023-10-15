package merge_sort

import (
	"fmt"
	"math"
)

func MergeSort(l []int) []int {

	if len(l) == 1 {
		return l
	}

	lLen := len(l)
	middle := math.Floor(float64(lLen) / 2)
	left := l[0:int(middle)]
	right := l[int(middle):lLen]
	fmt.Printf("left: %v\n", left)
	fmt.Printf("right: %v\n", right)
	return merge(
		MergeSort(left),
		MergeSort(right),
	)
}

func merge(right []int, left []int) []int {

	var res []int
	leftIndex := 0
	rightIndex := 0

	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			// push left
			res = append(res, left[leftIndex])
			leftIndex++
		} else {
			// push right
			res = append(res, right[rightIndex])
			rightIndex++
		}
	}

	res = append(res, left[leftIndex:]...)
	res = append(res, right[rightIndex:]...)

	return res
}
