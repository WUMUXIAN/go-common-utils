package trees

import (
	"errors"
	"testing"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearchTree(t *testing.T) {
	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
	})

	Convey("Delete Leaf Node Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{2, 1, 3}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		tree.Delete(1)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 3})
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2})
	})

	Convey("Delete Node Without Right Child Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{5, 3, 8, 2}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 5, 8})
		tree.Insert(6)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 5, 6})
	})

	Convey("Delete Node Without Left Child Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{5, 2, 3, 8}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		tree.Delete(2)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{3, 5, 8})
		tree.Insert(9)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{3, 5, 9})
	})

	Convey("Delete Node With Both Children Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{5, 3, 2, 4, 8, 6, 9}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 4, 5, 6, 8, 9})
		tree.Insert(7)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 4, 5, 6, 7, 9})
	})

	Convey("Inserting and Removing In A Row Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})

		tree.Delete(20)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 24, 25, 27, 30, 31, 35, 36})
		tree.Insert(20)

		tree.Delete(4)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
		tree.Delete(5)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})

		tree.Delete(16)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{9, 10, 13, 15, 20, 24, 25, 27, 30, 31, 35, 36})
		tree.Delete(15)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{9, 10, 13, 20, 24, 25, 27, 30, 31, 35, 36})
		tree.Delete(30)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{9, 10, 13, 20, 24, 25, 27, 31, 35, 36})
		tree.Delete(27)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{9, 10, 13, 20, 24, 25, 31, 35, 36})
		err := tree.Delete(100)
		So(err, ShouldBeError, errors.New("could not find node"))
	})

	Convey("Clear Tree Should Remove All Nodes From the Tree", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		unsorted := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		tree.Clear()
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{})
	})

}
