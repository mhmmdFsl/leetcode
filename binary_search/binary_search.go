package binary_search

import "math"

func BinarySearch(v int, l []int) bool {
	var arr []int
	if len(l) == 0 {
		return false
	}
	if len(l) == 1 {
		return l[0] == v
	}

	midleIndex := int(math.Floor(float64(len(l)) / 2))
	if l[midleIndex] == v {
		return true
	} else if l[midleIndex] < v {
		arr = l[midleIndex:len(l)]
	} else {
		arr = l[0:midleIndex]
	}

	return BinarySearch(v, arr)
}
