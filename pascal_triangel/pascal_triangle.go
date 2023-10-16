package pascal_triangel

func PascalTriangel(rowNum int) [][]int {

	var sliceOfSlice [][]int
	var lastSlice []int
	for i := 1; i <= rowNum; i++ {
		var currentSlice []int
		for j := 1; j <= i; j++ {
			if j == 1 || j == i {
				currentSlice = append(currentSlice, 1)
			} else {
				sum := lastSlice[j-2] + lastSlice[j-1]
				currentSlice = append(currentSlice, sum)
			}
		}
		lastSlice = currentSlice
		sliceOfSlice = append(sliceOfSlice, lastSlice)
	}
	return sliceOfSlice
}
