package trees

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSegmentTree(t *testing.T) {
	Convey("Minimum Range Query SegmentTree Should Work As Expected\n", t, func() {
		segmentTree := NewSegmentTree([]int{1, 5, 3, 7, 3, 2, 5, 7}, SegmentTreeMin)
		So(segmentTree.nodes, ShouldResemble, []int{0, 1, 1, 2, 1, 3, 2, 5, 1, 5, 3, 7, 3, 2, 5, 7})
		So(segmentTree.Query(1, 6), ShouldEqual, 2)
		So(segmentTree.Query(0, 4), ShouldEqual, 1)
		So(segmentTree.Query(3, 5), ShouldEqual, 2)
		So(segmentTree.Query(5, 3), ShouldEqual, -1)
		So(segmentTree.Query(-1, 3), ShouldEqual, -1)
		So(segmentTree.Query(3, 10), ShouldEqual, -1)
		segmentTree.Update(3, 9)
		So(segmentTree.nodes, ShouldResemble, []int{0, 1, 1, 2, 1, 3, 2, 5, 1, 5, 3, 9, 3, 2, 5, 7})
		segmentTree.Update(-1, 2)
		segmentTree.Update(10, 2)
		So(segmentTree.nodes, ShouldResemble, []int{0, 1, 1, 2, 1, 3, 2, 5, 1, 5, 3, 9, 3, 2, 5, 7})
	})
	Convey("Maximum Range Query SegmentTree Should Work As Expected", t, func() {
		segmentTree := NewSegmentTree([]int{1, 5, 3, 7, 3, 2, 5, 7}, SegmentTreeMax)
		So(segmentTree.nodes, ShouldResemble, []int{0, 7, 7, 7, 5, 7, 3, 7, 1, 5, 3, 7, 3, 2, 5, 7})
		So(segmentTree.Query(1, 6), ShouldEqual, 7)
		So(segmentTree.Query(0, 4), ShouldEqual, 7)
		So(segmentTree.Query(3, 5), ShouldEqual, 7)
		So(segmentTree.Query(5, 3), ShouldEqual, -1)
		So(segmentTree.Query(-1, 3), ShouldEqual, -1)
		So(segmentTree.Query(3, 10), ShouldEqual, -1)
		segmentTree.Update(3, 9)
		So(segmentTree.nodes, ShouldResemble, []int{0, 9, 9, 7, 5, 9, 3, 7, 1, 5, 3, 9, 3, 2, 5, 7})
		segmentTree.Update(-1, 2)
		segmentTree.Update(10, 2)
		So(segmentTree.nodes, ShouldResemble, []int{0, 9, 9, 7, 5, 9, 3, 7, 1, 5, 3, 9, 3, 2, 5, 7})
	})
	Convey("Sum Range Query SegmentTree Should Work As Expected", t, func() {
		segmentTree := NewSegmentTree([]int{1, 5, 3, 7, 3, 2, 5, 7}, SegmentTreeSum)
		So(segmentTree.nodes, ShouldResemble, []int{0, 33, 16, 17, 6, 10, 5, 12, 1, 5, 3, 7, 3, 2, 5, 7})
		So(segmentTree.Query(1, 6), ShouldEqual, 25)
		So(segmentTree.Query(0, 4), ShouldEqual, 19)
		So(segmentTree.Query(3, 5), ShouldEqual, 12)
		So(segmentTree.Query(5, 3), ShouldEqual, -1)
		So(segmentTree.Query(-1, 3), ShouldEqual, -1)
		So(segmentTree.Query(3, 10), ShouldEqual, -1)
		segmentTree.Update(3, 9)
		So(segmentTree.nodes, ShouldResemble, []int{0, 35, 18, 17, 6, 12, 5, 12, 1, 5, 3, 9, 3, 2, 5, 7})
		segmentTree.Update(-1, 2)
		segmentTree.Update(10, 2)
		So(segmentTree.nodes, ShouldResemble, []int{0, 35, 18, 17, 6, 12, 5, 12, 1, 5, 3, 9, 3, 2, 5, 7})

		segmentTree = NewSegmentTree([]int{-28, -39, 53, 65, 11, -56, -65, -39, -43, 97}, SegmentTreeSum)
		So(segmentTree.nodes, ShouldResemble, []int{0, -44, -117, 73, -50, -67, 118, -45, -104, 54, -28, -39, 53, 65, 11, -56, -65, -39, -43, 97})
		So(segmentTree.Query(1, 2), ShouldEqual, 14)
		So(segmentTree.Query(0, 4), ShouldEqual, 62)
	})
}
