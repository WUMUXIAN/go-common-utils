package trees

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
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

	Convey("Test With Large Amount Of Numbers Should Work", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		// Generate 10,000 unique numbers
		numbers := make(map[int]bool)
		for {
			if len(numbers) == 10000 {
				break
			}
			r := rand.Intn(500000)
			numbers[r] = true
		}
		reference := make([]int, 0)
		for value := range numbers {
			tree.Insert(value)
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

	Convey("Convert BST To Double Linked List Should Be OK", t, func() {
		tree := &BinarySearchTree{
			Root:       nil,
			Comparator: shared.IntComparator,
		}
		h, t := tree.ConvertToDoubleLinkedList()
		So(h, ShouldBeNil)
		So(t, ShouldBeNil)
		unsorted := []interface{}{20, 10, 30, 5, 15, 4, 9, 13, 16, 25, 35, 24, 27, 31, 36}
		for _, data := range unsorted {
			tree.Insert(data)
		}
		So(tree.ToSortedSlice(), ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})

		head, tail := tree.ConvertToDoubleLinkedList()

		headToTail := make([]interface{}, 0)
		tailToHead := []interface{}{}
		for {
			if head == nil {
				break
			}
			headToTail = append(headToTail, head.Data)
			head = head.Right
		}

		for {
			if tail == nil {
				break
			}
			tailToHead = append(tailToHead, tail.Data)
			tail = tail.Left
		}
		So(headToTail, ShouldResemble, []interface{}{4, 5, 9, 10, 13, 15, 16, 20, 24, 25, 27, 30, 31, 35, 36})
		So(tailToHead, ShouldResemble, []interface{}{36, 35, 31, 30, 27, 25, 24, 20, 16, 15, 13, 10, 9, 5, 4})
	})

}
