package trygo

// Append func
func Append(slice []int, element int) []int {
	n := len(slice)
	if cap(slice) == n {
		slice = DoubleExtend(slice)
	}
	slice = slice[0 : n+1]
	slice[n] = element
	return slice
}

// DoubleExtend func
func DoubleExtend(slice []int) []int {
	var newcap int
	if cap(slice) == 0 {
		newcap = 16
	} else {
		newcap = 2 * cap(slice)
	}
	newSlice := make([]int, len(slice), newcap)

	for i := range slice {
		newSlice[i] = slice[i]
	}
	return newSlice
}
