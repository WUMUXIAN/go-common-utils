package trees

import (
	"testing"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearchTree(t *testing.T) {
	tree := &BinarySearchTree{
		Root:       nil,
		Comparator: shared.IntComparator,
	}

	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(4)
	tree.Insert(9)
	tree.Insert(13)
	tree.Insert(16)
	tree.Insert(25)
	tree.Insert(35)
	tree.Insert(24)
	tree.Insert(27)
	tree.Insert(31)
	tree.Insert(36)

	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
	})

	Convey("Delete Should Work", t, func() {
		tree.Delete(4)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
		tree.Delete(16)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{5, 9, 10, 13, 15, 20, 24, 25, 27, 30, 31, 35, 36})
		tree.Delete(25)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{5, 9, 10, 13, 15, 20, 24, 27, 30, 31, 35, 36})
		tree.Delete(20)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{5, 9, 10, 13, 15, 24, 27, 30, 31, 35, 36})
	})

	Convey("Clear Tree Should Remove All Nodes From the Tree", t, func() {
		tree.Clear()
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{})
	})

}
