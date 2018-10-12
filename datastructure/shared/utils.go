package shared

// BinarySearch performs binary search for a value in sorted value slice.
func BinarySearch(sortedValues []interface{}, searchValue interface{}, comparator Comparator) int {
	if len(sortedValues) == 0 {
		return -1
	}
	start := 0
	end := len(sortedValues) - 1
	for start <= end {
		mid := (start + end) / 2
		if comparator(sortedValues[mid], searchValue) == 0 {
			return mid
		} else if comparator(sortedValues[mid], searchValue) > 0 {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return -1
}
