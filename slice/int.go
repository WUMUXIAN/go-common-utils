package slice

import (
	"fmt"
	"math"
	"math/rand"
)

// IndexOfInt gets the index of an int element in an int slice
func IndexOfInt(x []int, y int) int {
	for i, v := range x {
		if v == y {
			return i
		}
	}
	return -1
}

// ContainsInt checks whether an int element is in an int slice
func ContainsInt(x []int, y int) bool {
	return IndexOfInt(x, y) != -1
}

// EqualsInts checks whether two int slice has the same elements
func EqualsInts(x []int, y []int) bool {
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

// CopyInts makes a new int slice that copies the content of the given int slice
func CopyInts(x []int) []int {
	return append([]int{}, x...)
}

// CutInts cuts an int slice by removing the elements starts from i and ends at j-1
func CutInts(x []int, i, j int) ([]int, error) {
	if i < 0 || j > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	if i >= j {
		return x, fmt.Errorf("%d must be smaller than %d", i, j)
	}
	return append(x[:i], x[j:]...), nil
}

// RemoveInt removes an int from a given int slice by value
func RemoveInt(x []int, y int) []int {
	index := IndexOfInt(x, y)
	if index != -1 {
		return append(x[:index], x[(index+1):]...)
	}
	return x
}

// RemoveIntAt removes an int from a given int slice by index
func RemoveIntAt(x []int, index int) ([]int, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], x[(index+1):]...), nil
}

// InsertIntAt inserts an int value into a given int slice at given index
func InsertIntAt(x []int, y int, index int) ([]int, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	x = append(x, 0)
	copy(x[index+1:], x[index:])
	x[index] = y
	return x, nil
}

// InsertIntsAt inserts a int slice into a given int slice at given index
func InsertIntsAt(x []int, y []int, index int) ([]int, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], append(y, x[index:]...)...), nil
}

// PopFirstInt pops the first value of an int slice
func PopFirstInt(x []int) (int, []int, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[0], x[1:], nil
}

// PopLastInt pops the last value of an int slice
func PopLastInt(x []int) (int, []int, error) {
	if len(x) == 0 {
		return 0, nil, fmt.Errorf("no value to pop")
	}
	return x[len(x)-1], x[:len(x)-1], nil
}

// FilterInts filters an int slice by the given filter function
func FilterInts(x []int, filter func(int) bool) []int {
	y := x[:0]
	for _, v := range x {
		if filter(v) {
			y = append(y, v)
		}
	}
	return y
}

// ReverseInts reverses an int slice
func ReverseInts(x []int) []int {
	for i := len(x)/2 - 1; i >= 0; i-- {
		opp := len(x) - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

// ShuffleInts shuffles an int slice
func ShuffleInts(x []int) []int {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// MergeInts merges two int slice with specific excluded values
func MergeInts(x []int, y []int, excludes ...int) []int {
	traceMap := make(map[int]bool)
	result := make([]int, 0)
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

// SumOfInts find the sum of all items in int slice
func SumOfInts(x []int) int {
	var sum = 0
	for _, v := range x {
		sum += v
	}
	return sum
}

// TransformInts helps figure out how to transform current to target slice by returning the ones to add and remove
func TransformInts(target, current []int) (add, remove []int) {
	add = make([]int, 0)
	remove = make([]int, 0)

	// Process
	if target != nil {
		statusMap := make(map[int]int) // the int is the status, -1: to be removed, 0: stay there, 1: to be added.
		length := int(math.Max(float64(len(target)), float64(len(current))))
		for i := 0; i < length; i++ {
			if i <= len(target)-1 {
				if _, ok := statusMap[target[i]]; ok {
					statusMap[target[i]]++
				} else {
					statusMap[target[i]] = 1
				}
			}
			if i <= len(current)-1 {
				if _, ok := statusMap[current[i]]; ok {
					statusMap[current[i]]--
				} else {
					statusMap[current[i]] = -1
				}
			}
		}
		for v, status := range statusMap {
			if status < 0 {
				remove = append(remove, v)
			} else if status > 0 {
				add = append(add, v)
			}
		}
	}

	return
}
