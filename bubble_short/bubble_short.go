package bubbleshort

import "fmt"

func BubbleShort(list []int) []int {

	for i := 0; i < len(list)-1; i++ {
		for j := 1; j <= len(list)-1; j++ {
			var temp int
			if list[j-1] > list[j] {
				temp = list[j-1]
				list[j-1] = list[j]
				list[j] = temp
			}
		}
	}

	fmt.Printf("%v\n", list)

	return list
}
