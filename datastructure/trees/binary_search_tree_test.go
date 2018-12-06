package trees

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
	"testing"

	"github.com/TectusDreamlab/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBinarySearchTree(t *testing.T) {
	Convey("Inorder Traverse Should Generate Sorted Values", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		unsortedValues := []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o"}
		for i, data := range unsortedKeys {
			tree.Put(data, unsortedValues[i])
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
		So(tree.GetSize(), ShouldEqual, 15)

		So(tree.BreadthFirstTraverse(), ShouldResemble, []interface{}{20, 10, 30, 5, 15, 25, 35, 4, 9, 13, 16, 24, 27, 31, 36})
		So(tree.DepthFirstTraverse(), ShouldResemble, []interface{}{20, 10, 5, 4, 9, 15, 13, 16, 30, 25, 24, 27, 35, 31, 36})

		So(tree.Search(35), ShouldEqual, "k")
		So(tree.Search(9), ShouldEqual, "g")
		So(tree.Search(37), ShouldBeNil)
	})

	Convey("Put Value For Existing Key", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{2, 1, 3}
		unsortedValues := []interface{}{"a", "b", "c"}
		for i, data := range unsortedKeys {
			tree.Put(data, unsortedValues[i])
		}
		So(tree.Search(1), ShouldEqual, "b")
		tree.Put(1, "d")
		So(tree.Search(1), ShouldEqual, "d")
		So(tree.GetSize(), ShouldEqual, 3)
		tree.Put(4, "e")
		tree.Put(2, "f")
		So(tree.GetSize(), ShouldEqual, 4)
	})

	Convey("Delete Leaf Node Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{2, 1, 3}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		tree.Delete(1)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 3})
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2})
	})

	Convey("Delete Node Without Right Child Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{5, 3, 8, 2}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 5, 8})
		tree.Put(6, nil)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 5, 6})
	})

	Convey("Delete Node Without Left Child Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{5, 2, 3, 8}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		tree.Delete(2)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{3, 5, 8})
		tree.Put(9, nil)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{3, 5, 9})
	})

	Convey("Delete Node With Both Children Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{5, 3, 2, 4, 8, 6, 9}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		tree.Delete(3)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 4, 5, 6, 8, 9})
		tree.Put(7, nil)
		tree.Delete(8)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{2, 4, 5, 6, 7, 9})
	})

	Convey("Inserting and Removing In A Row Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})

		tree.Delete(20)
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 24, 25, 27, 30, 31, 35, 36})
		tree.Put(20, nil)

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
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		unsortedKeys := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		tree.Clear()
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{})
	})

	Convey("Test With Large Amount Of Numbers Should Work", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		// Generate 10,000 unique numbers
		numbers := make(map[int]bool)
		for len(numbers) < 10000 {
			r := rand.Intn(500000)
			numbers[r] = true
		}
		reference := make([]int, 0)
		for value := range numbers {
			tree.Put(value, nil)
			reference = append(reference, value)
		}
		sorted := tree.ToSortedSlice()
		sort.Ints(reference)
		for i := range sorted {
			if sorted[i].(int) != reference[i] {
				So(false, ShouldBeTrue)
			}
		}

		// Let's remove some fields.
		tree.Delete(reference[10000/2])
		sorted = tree.ToSortedSlice()
		reference = append(reference[:10000/2], reference[10000/2+1:]...)
		for i := range sorted {
			if sorted[i].(int) != reference[i] {
				fmt.Println(i, sorted[i], reference[i])
				So(false, ShouldBeTrue)
			}
		}

		tree.Delete(reference[10000/4])
		sorted = tree.ToSortedSlice()
		reference = append(reference[:10000/4], reference[10000/4+1:]...)
		for i := range sorted {
			if sorted[i].(int) != reference[i] {
				fmt.Println(sorted[i], reference[i])
				So(false, ShouldBeTrue)
			}
		}

		tree.Delete(reference[10000/8])
		sorted = tree.ToSortedSlice()
		reference = append(reference[:10000/8], reference[10000/8+1:]...)
		for i := range sorted {
			if sorted[i].(int) != reference[i] {
				fmt.Println(sorted[i], reference[i])
				So(false, ShouldBeTrue)
			}
		}
	})

	Convey("Convert BST To Double Linked List And Convert Back Should Be OK", t, func() {
		tree := &BinarySearchTree{
			BinaryTree: BinaryTree{
				Root:       nil,
				Comparator: shared.IntComparator,
			},
		}
		h, t := tree.ConvertToDoubleLinkedList()
		So(h, ShouldBeNil)
		So(t, ShouldBeNil)
		unsortedKeys := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsortedKeys {
			tree.Put(data, nil)
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})

		head, tail := tree.ConvertToDoubleLinkedList()

		headToTail := make([]interface{}, 0)
		tailToHead := []interface{}{}
		p := head
		for p != nil {
			headToTail = append(headToTail, p.Key)
			p = p.Right
		}

		p = tail
		for p != nil {
			tailToHead = append(tailToHead, p.Key)
			p = p.Left
		}
		So(headToTail, ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
		So(tailToHead, ShouldResemble, []interface{}{36, 35, 31, 30, 27, 25, 24, 20, 16, 15, 13, 10, 9, 5, 4})

		tree.ConvertFromDoubleLinkedList(head, tree.GetSize())
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
	})

}
