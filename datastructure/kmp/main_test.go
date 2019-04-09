// Package kmp implements the KMP algorithm
package kmp

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSearchOccurrence(t *testing.T) {
	testCases := []struct {
		inputs []string
		output []int
	}{
		{[]string{"abababc", "ab"}, []int{0, 2, 4}},
		{[]string{"abcabcababc", "abcabcaba"}, []int{0}},
		{[]string{"abcabcababc", "bca"}, []int{1, 4}},
		{[]string{"what are you doing are?", "are"}, []int{5, 19}},
		{[]string{"abcabcababc", "abcabcabc"}, nil},
		{[]string{"", "abcabcabc"}, nil},
		{[]string{"abcabcababc", ""}, nil},
		{[]string{"bbc", "d"}, nil},
		{[]string{"bbc", "abcdaa"}, nil},
		{[]string{"abcefdgdafawersf", "fawer"}, []int{9}},
	}
	for _, testCase := range testCases {
		Convey(fmt.Sprintf("Search %s In %s Should Return Occurrence %v\n", testCase.inputs[1], testCase.inputs[0], testCase.output), t, func() {
			So(SearchOccurrence(testCase.inputs[0], testCase.inputs[1]), ShouldResemble, testCase.output)
		})
	}
}
