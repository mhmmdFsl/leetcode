package linear_search

import "fmt"

func LinearSearch(v int, arr []int) (int, error) {

	for i := range arr {
		if arr[i] == v {
			return i, nil
		}
	}

	return 0, fmt.Errorf("value not found")
}
