package golib

// FindIndexSlice returns the index of a value in a slice
func FindIndexSlice(slice []int, value int) int {
	for p, v := range slice {
		if value == v {
			return p
		}
	}
	return -1
}
