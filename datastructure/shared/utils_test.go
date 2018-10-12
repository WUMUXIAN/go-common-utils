package shared

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	sortedValues := []interface{}{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		So(BinarySearch(sortedValues, 3, IntComparator), ShouldEqual, 1)
		So(BinarySearch(sortedValues, 9, IntComparator), ShouldEqual, 4)
		So(BinarySearch(sortedValues, 5, IntComparator), ShouldEqual, 2)
		So(BinarySearch(sortedValues, 6, IntComparator), ShouldEqual, 2)
		So(BinarySearch(sortedValues, 16, IntComparator), ShouldEqual, 7)
		So(BinarySearch(sortedValues, 0, IntComparator), ShouldEqual, -1)
		So(BinarySearch(sortedValues, 20, IntComparator), ShouldEqual, 9)
	})
}
