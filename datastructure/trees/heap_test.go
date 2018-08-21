package trees

import (
	"testing"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestHeap(t *testing.T) {
	Convey("Build Max Heap With With A List Of Values Should Be OK", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMax,
			Comparator: shared.IntComparator,
		}
		So(heap.GetValues(), ShouldResemble, []interface{}{})
		values := []interface{}{1, 8, 10, 2, 11, 7, 9, 3, 4, 5, 6, 20}
		heap.InitHeap(values)
		So(heap.GetValues(), ShouldResemble, []interface{}{20, 11, 10, 4, 8, 7, 9, 3, 2, 5, 6, 1})
	})

	Convey("Insert Node Into Max Heap Should Be OK", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMax,
			Comparator: shared.IntComparator,
		}
		heap.Insert(1)
		So(heap.GetValues(), ShouldResemble, []interface{}{1})
		heap.Insert(8)
		So(heap.GetValues(), ShouldResemble, []interface{}{8, 1})
	})

	Convey("bubbleup Should Not Do Anything When Heap Size Is 1", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMax,
			Comparator: shared.IntComparator,
		}
		heap.Insert(1)
		heap.bubbleUp()
		So(heap.GetValues(), ShouldResemble, []interface{}{1})
	})

	Convey("Peek And Pop From Max Heap Should Work", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMax,
			Comparator: shared.IntComparator,
		}
		So(heap.GetValues(), ShouldResemble, []interface{}{})
		So(heap.Peek(), ShouldEqual, nil)
		So(heap.Pop(), ShouldEqual, nil)
		values := []interface{}{1, 8, 10, 2, 11, 7, 9, 3, 4, 5, 6, 20}
		heap.InitHeap(values)
		var top interface{}
		top = heap.Pop()
		So(top, ShouldEqual, 20)
		top = heap.Peek()
		So(top, ShouldEqual, 11)
		heap.Insert(50)
		top = heap.Peek()
		So(top, ShouldEqual, 50)
		heap.Insert(1)
		So(heap.GetValues(), ShouldResemble, []interface{}{50, 8, 11, 4, 6, 10, 9, 3, 2, 5, 1, 7, 1})
		top = heap.Pop()
		So(top, ShouldEqual, 50)
		So(heap.GetValues(), ShouldResemble, []interface{}{11, 8, 10, 4, 6, 7, 9, 3, 2, 5, 1, 1})
	})

	Convey("Build Min Heap With With A List Of Values Should Be OK", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMin,
			Comparator: shared.IntComparator,
		}
		So(heap.GetValues(), ShouldResemble, []interface{}{})
		values := []interface{}{1, 8, 10, 2, 11, 7, 9, 3, 4, 5, 6, 20}
		heap.InitHeap(values)
		So(heap.GetValues(), ShouldResemble, []interface{}{1, 2, 7, 3, 5, 10, 9, 8, 4, 11, 6, 20})
	})

	Convey("Insert Node Into Min Heap Should Be OK", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMin,
			Comparator: shared.IntComparator,
		}
		heap.Insert(8)
		So(heap.GetValues(), ShouldResemble, []interface{}{8})
		heap.Insert(1)
		So(heap.GetValues(), ShouldResemble, []interface{}{1, 8})
	})

	Convey("Peek And Pop From Min Heap Should Work", t, func() {
		heap := &Heap{
			HeapType:   HeapTypeMin,
			Comparator: shared.IntComparator,
		}
		So(heap.GetValues(), ShouldResemble, []interface{}{})
		So(heap.Peek(), ShouldEqual, nil)
		So(heap.Pop(), ShouldEqual, nil)
		values := []interface{}{1, 8, 10, 2, 11, 7, 9, 3, 4, 5, 6, 20}
		heap.InitHeap(values)
		var top interface{}
		top = heap.Pop()
		So(top, ShouldEqual, 1)
		top = heap.Peek()
		So(top, ShouldEqual, 2)
		heap.Insert(1)
		top = heap.Peek()
		So(top, ShouldEqual, 1)
		top = heap.Pop()
		top = heap.Pop()
		top = heap.Pop()
		So(top, ShouldEqual, 3)
	})
}
