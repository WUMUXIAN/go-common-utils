// Package kmp implements the KMP algorithm
package kmp

import "fmt"

// SearchOccurrence search the occurrences of pattern in str.
// If pattern never occur in str, return empty string.
// The time complexity for KMP is O(N), where N is the length of str.
func SearchOccurrence(str, pattern string) (indices []int) {
	// if str is empty, no occurrence is possible.
	// if pattern is empty, no occurrence is existed.
	if len(str) == 0 || len(pattern) == 0 {
		return nil
	}

	// When using KMP algorithm, the first thing we do is to construct the next array for pattern.
	// note taht next[0] = 0 and will never be used.
	next := make([]int, len(pattern))
	
	// To calculate the next array, we try to match the pattern string against itself, starting from the second character.

	j := 0
	for i := 1; i < len(pattern); i++ {
		// if j is not 0 and we find a mismatch, move j to next[j] until we find a match or j goes back to 0
		for j > 0 && pattern[i] != pattern[j] {
			j = next[j]
		}
		// if we find a match here, then it means for i+1, the next is j+1.
		// if i + 1 >= pattern, we don't have to track it.
		if pattern[i] == pattern[j] {
			if i+1 < len(pattern) {
				next[i+1] = j + 1
			}
			j++
		}
		// if j == 0, means for i+1, the next is 0, since 0 is the default value, we don't have to do anything.
	}
	// Now we have the next slice constructed, we can start the search process.
	j = 0
	i := 0
	for i < len(str) {
		if str[i] != pattern[j] {
			fmt.Printf("we encounter a mismatch, str[%d] is %c, pattern[%d] is %c\n", i, str[i], j, pattern[j])
			// when we have a mismatch and j is not zero, we keep i as it is and move j backword to next[j]
			if j > 0 {
				j = next[j]
			} else {
				// if j is zero, means we have exausted all possibilities, move i forward now.
				i++
			}
		} else {
			// when we have a match, move i and j forward together.
			i++
			j++

			// if j is at the end, then we've found an occurrence. record the starting index (i-j) and reset j to 0 to start looking again.
			if j == len(pattern) {
				indices = append(indices, i-j)
				j = 0
			}
		}
	}
	return indices
}
