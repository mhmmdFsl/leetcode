package selectionshort

import "fmt"

func SelectionSort(list []int) []int {

	for i := range list {
		currentIndex := i
		smallestValue := list[currentIndex]
		smallestValueIndex := i
		for j := currentIndex + 1; j < len(list); j++ {
			if list[j] < smallestValue {
				smallestValue = list[j]
				smallestValueIndex = j
			}
		}

		temp := list[currentIndex]
		list[currentIndex] = smallestValue
		list[smallestValueIndex] = temp

	}

	fmt.Printf("%v\n", list)

	return list
}
