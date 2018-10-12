package shared

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	sortedValues := []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		So(BinarySearch(sortedValues, 3, IntComparator), ShouldEqual, 2)
		So(BinarySearch(sortedValues, 9, IntComparator), ShouldEqual, 8)
		So(BinarySearch(sortedValues, 5, IntComparator), ShouldEqual, 4)
		So(BinarySearch(sortedValues, 0, IntComparator), ShouldEqual, -1)
		So(BinarySearch(sortedValues, 11, IntComparator), ShouldEqual, -1)
	})
}
