// Package slice wraps up some common operations regarding to golang slice with different types.
package slice

import (
	"fmt"
	"math"
	"math/rand"
)

// IndexOfString gets the index of a string element in a string slice
func IndexOfString(x []string, y string) int {
	for i, v := range x {
		if v == y {
			return i
		}
	}
	return -1
}

// ContainsString checks whether a string element is in a string slice
func ContainsString(x []string, y string) bool {
	return IndexOfString(x, y) != -1
}

// EqualsStrings checks whether two string slice has the same elements
func EqualsStrings(x []string, y []string) bool {
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

// CopyStrings makes a new string slice that copies the content of the given string slice
func CopyStrings(x []string) []string {
	return append([]string{}, x...)
}

// CutStrings cuts a string slice by removing the elements starts from i and ends at j-1
func CutStrings(x []string, i, j int) ([]string, error) {
	if i < 0 || j > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	if i >= j {
		return x, fmt.Errorf("%d must be smaller than %d", i, j)
	}
	return append(x[:i], x[j:]...), nil
}

// RemoveString removes a string from a given string slice by value
func RemoveString(x []string, y string) []string {
	index := IndexOfString(x, y)
	if index != -1 {
		return append(x[:index], x[(index+1):]...)
	}
	return x
}

// RemoveStringAt removes a string from a given string slice by index
func RemoveStringAt(x []string, index int) ([]string, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], x[(index+1):]...), nil
}

// InsertStringAt inserts a string value stringo a given string slice at given index
func InsertStringAt(x []string, y string, index int) ([]string, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	x = append(x, "")
	copy(x[index+1:], x[index:])
	x[index] = y
	return x, nil
}

// InsertStringsAt inserts a string slice stringo a given string slice at given index
func InsertStringsAt(x []string, y []string, index int) ([]string, error) {
	if index < 0 || index > len(x) {
		return x, fmt.Errorf("out of bound")
	}
	return append(x[:index], append(y, x[index:]...)...), nil
}

// PopFirstString pops the first value of a string slice
func PopFirstString(x []string) (string, []string, error) {
	if len(x) == 0 {
		return "", nil, fmt.Errorf("no value to pop")
	}
	return x[0], x[1:], nil
}

// PopLastString pops the last value of a string slice
func PopLastString(x []string) (string, []string, error) {
	if len(x) == 0 {
		return "", nil, fmt.Errorf("no value to pop")
	}
	return x[len(x)-1], x[:len(x)-1], nil
}

// FilterStrings filters a string slice by the given filter function
func FilterStrings(x []string, filter func(string) bool) []string {
	y := x[:0]
	for _, v := range x {
		if filter(v) {
			y = append(y, v)
		}
	}
	return y
}

// ReverseStrings reverses a string slice
func ReverseStrings(x []string) []string {
	for i := len(x)/2 - 1; i >= 0; i-- {
		opp := len(x) - 1 - i
		x[i], x[opp] = x[opp], x[i]
	}
	return x
}

// ShuffleStrings shuffles a string slice
func ShuffleStrings(x []string) []string {
	for i := len(x) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		x[i], x[j] = x[j], x[i]
	}
	return x
}

// MergeStrings merges two string slice with specific excluded values
func MergeStrings(x []string, y []string, excludes ...string) []string {
	traceMap := make(map[string]bool)
	result := make([]string, 0)
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

// UniqueStrings removes the duplicates from the string slice
func UniqueStrings(x []string) []string {
	traceMap := make(map[string]bool)
	result := make([]string, 0)
	for _, v := range x {
		if _, value := traceMap[v]; !value {
			traceMap[v] = true
			result = append(result, v)
		}
	}
	return result
}

// TransformStrings helps figure out how to transform current to target slice by returning the ones to add and remove
func TransformStrings(target, current []string) (add, remove []string) {
	add = make([]string, 0)
	remove = make([]string, 0)

	// Process
	statusMap := make(map[string]int) // the int is the status, -1: to be removed, 0: stay there, 1: to be added.
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

	return
}
