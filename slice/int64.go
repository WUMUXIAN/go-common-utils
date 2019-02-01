package slice

import (
	"fmt"
	"math/rand"
)

// IndexOfInt64 gets the index of an int64 element in an int64 slice
func IndexOfInt64(x []int64, y int64) int {
	for i, v := range x {
		if v == y {
			return i
		}
	}
	return -1
}

// ContainsInt64 checks whether an int64 element is in an int64 slice
func ContainsInt64(x []int64, y int64) bool {
	return IndexOfInt64(x, y) != -1
}

// EqualsInt64s checks whether two int64 slice has the same elements
func EqualsInt64s(x []int64, y []int64) bool {
	if len(x) != len(y) {
		return false
	}

	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

// CopyInt64s makes a new int64 slice that copies the content of the given int64 slice
func CopyInt64s(x []int64) []int64 {
	return append([]int64{}, x...)
}

// CutInt64s cuts an int64 slice by removing the elements starts from i and ends at j-1
func CutInt64s(x []int64, i, j int) ([]int64, error) {
	if i < 0 || j > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	if i >= j {
		return x, fmt.Errorf("%d must be smaller than %d", i, j)
	}
	return append(x[:i], x[j:]...), nil
}

// RemoveInt64 removes an int64 from a given int64 slice by value
func RemoveInt64(x []int64, y int64) []int64 {
	index := IndexOfInt64(x, y)
	if index != -1 {
		return append(x[:index], x[(index+1):]...)
	}
	return x
}

// RemoveInt64At removes an int64 from a given int64 slice by index
func RemoveInt64At(x []int64, index int) ([]int64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], x[(index+1):]...), nil
}

// InsertInt64At inserts an int64 value into a given int64 slice at given index
func InsertInt64At(x []int64, y int64, index int) ([]int64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	x = append(x, 0)
	copy(x[index+1:], x[index:])
	x[index] = y
	return x, nil
}

// InsertInt64sAt inserts a int64 slice into a given int64 slice at given index
func InsertInt64sAt(x []int64, y []int64, index int) ([]int64, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], append(y, x[index:]...)...), nil
}

// PopFirstInt64 pops the first value of an int64 slice
func PopFirstInt64(x []int64) (int64, []int64, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[0], x[1:], nil
}

// PopLastInt64 pops the last value of an int64 slice
func PopLastInt64(x []int64) (int64, []int64, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[len(x)-1], x[:len(x)-1], nil
}

// FilterInt64s filters an int64 slice by the given filter function
func FilterInt64s(x []int64, filter func(int64) bool) []int64 {
	y := x[:0]
	for _, v := range x {
		if filter(v) {
			y = append(y, v)
		}
	}
	return y
}

// ReverseInt64s reverses an int64 slice
func ReverseInt64s(x []int64) []int64 {
	for i := len(x)/2 - 1; i >= 0; i-- {
		opp := len(x) - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

// ShuffleInt64s shuffles an int64 slice
func ShuffleInt64s(x []int64) []int64 {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// MergeInt64s merges two int64 slice with specific excluded values
func MergeInt64s(x []int64, y []int64, excludes ...int64) []int64 {
	traceMap := make(map[int64]bool)
	result := make([]int64, 0)
	for _, ex := range excludes {
		traceMap[ex] = true
	}
	// We preserve the order by x and then y
	for _, v := range x {
		if !traceMap[v] {
			traceMap[v] = true
			result = append(result, v)
		}
	}

	for _, v := range y {
		if !traceMap[v] {
			traceMap[v] = true
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt64s returns the intersection of two int64 slices
func IntersectInt64s(x []int64, y []int64) []int64 {
	traceMap := make(map[int64]bool)
	result := make([]int64, 0)
	for _, v := range x {
		traceMap[v] = true
	}
	for _, v := range y {
		if traceMap[v] {
			result = append(result, v)
		}
	}
	return result
}
