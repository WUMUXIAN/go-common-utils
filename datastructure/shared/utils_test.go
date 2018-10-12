package shared

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearch(t *testing.T) {
	sortedValues := []interface{}{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		var index int
		var found bool
		index, found = BinarySearch(sortedValues, 3, IntComparator)
		So(index, ShouldEqual, 1)
		So(found, ShouldEqual, true)
		index, found = BinarySearch(sortedValues, 9, IntComparator)
		So(index, ShouldEqual, 4)
		So(found, ShouldEqual, true)
		index, found = BinarySearch(sortedValues, 5, IntComparator)
		So(index, ShouldEqual, 2)
		So(found, ShouldEqual, true)
		index, found = BinarySearch(sortedValues, 6, IntComparator)
		So(index, ShouldEqual, 2)
		So(found, ShouldEqual, false)
		index, found = BinarySearch(sortedValues, 16, IntComparator)
		So(index, ShouldEqual, 7)
		So(found, ShouldEqual, false)
		index, found = BinarySearch(sortedValues, 0, IntComparator)
		So(index, ShouldEqual, -1)
		So(found, ShouldEqual, false)
		index, found = BinarySearch(sortedValues, 20, IntComparator)
		So(index, ShouldEqual, 9)
		So(found, ShouldEqual, false)
	})
}
