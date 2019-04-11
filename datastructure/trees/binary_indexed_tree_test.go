package trees

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestBinaryIndexedTree(t *testing.T) {
	Convey("Binary Indexed Tree Should Work As Expected\n", t, func() {
		bit := NewBinaryIndexedTree([]int{1, 5, 3, 7, 3, 2, 5, 7})
		So(bit.c, ShouldResemble, []int{0, 1, 6, 3, 16, 3, 5, 5, 33})
		So(bit.GetSum(4), ShouldEqual, 19)
		So(bit.RangeSum(4, 6), ShouldEqual, 10)
		So(bit.RangeSum(0, 3), ShouldEqual, 16)
		So(bit.RangeSum(-1, 3), ShouldEqual, -1)
		So(bit.RangeSum(3, 10), ShouldEqual, -1)
		So(bit.GetSum(-1), ShouldEqual, -1)
		So(bit.GetSum(10), ShouldEqual, -1)

		bit.Update(0, 2)
		So(bit.c, ShouldResemble, []int{0, 3, 8, 3, 18, 3, 5, 5, 35})
		bit.Update(4, 2)
		So(bit.c, ShouldResemble, []int{0, 3, 8, 3, 18, 5, 7, 5, 37})
		So(bit.RangeSum(0, 7), ShouldEqual, 37)
	})
}
