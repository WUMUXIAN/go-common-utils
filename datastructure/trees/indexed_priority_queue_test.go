package trees

import (
	"errors"
	"testing"

	"github.com/WUMUXIAN/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestIndexedPriorityQueue(t *testing.T) {
	indexedPriorityQueue := NewIndexedPriorityQueue(10, HeapTypeMin, shared.StringComparator)
	Convey("Create An New Indexed Priority Queue", t, func() {
		Convey("Should Be Empty", func() {
			So(indexedPriorityQueue.IsEmpty(), ShouldBeTrue)
		})
		Convey("Size Should Be Zero", func() {
			So(indexedPriorityQueue.Size(), ShouldEqual, 0)
		})
		Convey("Peek And Pop Should Return Error", func() {
			_, _, err := indexedPriorityQueue.Peek()
			So(err, ShouldBeError, errors.New("queue is empty"))
			_, _, err = indexedPriorityQueue.Pop()
			So(err, ShouldBeError, errors.New("queue is empty"))
		})
		Convey("It Should Not Contain Any Value At An Valid Index", func() {
			So(indexedPriorityQueue.Contains(2), ShouldBeFalse)
		})
	})

	strs := []string{"it", "was", "the", "best", "of", "times", "it", "was", "the", "worst"}
	Convey("Inserting 10 Values Into The Priority", t, func() {
		for i, str := range strs {
			indexedPriorityQueue.Insert(i, str)
		}
		Convey("The Size Should Be 10", func() {
			So(indexedPriorityQueue.Size(), ShouldEqual, 10)
		})
		Convey("Reinsert At The Same Index Should Fail", func() {
			err := indexedPriorityQueue.Insert(0, "abc")
			So(err, ShouldBeError, errors.New("index is already used"))
		})
		Convey("Reinsert At The Invalid Index Should Fail", func() {
			err := indexedPriorityQueue.Insert(-1, "abc")
			So(err, ShouldBeError, errors.New("index out of range"))
			err = indexedPriorityQueue.Insert(11, "abc")
			So(err, ShouldBeError, errors.New("index out of range"))
		})
	})

	Convey("Peek And Pop The Queue", t, func() {
		targetOrder := []int{3, 6, 0, 4, 2, 8, 5, 7, 1, 9}
		output := ""
		for i := range targetOrder {
			output += " " + strs[targetOrder[i]]
		}
		Convey("Peek And Pop Should Work, The Target Order Should Be --->"+output, func() {
			i := 0
			for !indexedPriorityQueue.IsEmpty() {
				index, value, err := indexedPriorityQueue.Peek()
				So(err, ShouldBeNil)
				So(index, ShouldEqual, targetOrder[i])
				So(value, ShouldEqual, strs[targetOrder[i]])
				index, value, err = indexedPriorityQueue.Pop()
				So(err, ShouldBeNil)
				So(index, ShouldEqual, targetOrder[i])
				So(value, ShouldEqual, strs[targetOrder[i]])
				i++
			}
			So(indexedPriorityQueue.Size(), ShouldEqual, 0)
		})
	})

	Convey("Get Value At Given Index", t, func() {
		for i, str := range strs {
			indexedPriorityQueue.Insert(i, str)
		}
		Convey("Get By Invalid Index Should Fail", func() {
			_, err := indexedPriorityQueue.GetValue(-1)
			So(err, ShouldBeError, errors.New("index out of range"))
			_, err = indexedPriorityQueue.GetValue(11)
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("Get By Correct Index Should Work", func() {
			value, err := indexedPriorityQueue.GetValue(5)
			So(err, ShouldBeNil)
			So(value, ShouldEqual, "times")
			value, err = indexedPriorityQueue.GetValue(7)
			So(err, ShouldBeNil)
			So(value, ShouldEqual, "was")
		})
		indexedPriorityQueue.Pop()
		Convey("Get By Index That Contains No Value Should Fail", func() {
			_, err := indexedPriorityQueue.GetValue(3)
			So(err, ShouldBeError, errors.New("index does not have value"))
		})
	})

	Convey("Change Value At Given Index", t, func() {
		Convey("Change By Invalid Index Should Fail", func() {
			err := indexedPriorityQueue.ChangeValue(-1, "am")
			So(err, ShouldBeError, errors.New("index out of range"))
			err = indexedPriorityQueue.ChangeValue(11, "am")
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("Change By Index That Contains No Value Should Fail", func() {
			err := indexedPriorityQueue.ChangeValue(3, "am")
			So(err, ShouldBeError, errors.New("index does not have value"))
		})
		Convey("Change By Correct Index Should Work", func() {
			err := indexedPriorityQueue.ChangeValue(9, "am")
			So(err, ShouldBeNil)
			index, value, err := indexedPriorityQueue.Peek()
			So(err, ShouldBeNil)
			So(index, ShouldEqual, 9)
			So(value, ShouldEqual, "am")
			So(indexedPriorityQueue.pq[3], ShouldEqual, 6)
			So(indexedPriorityQueue.pq[7], ShouldEqual, 2)

			err = indexedPriorityQueue.ChangeValue(0, "yeild")
			So(indexedPriorityQueue.pq[2], ShouldEqual, 4)
			So(indexedPriorityQueue.pq[5], ShouldEqual, 0)
		})
	})

	Convey("Delete Value At Given Index", t, func() {
		Convey("Delete By Invalid Index Should Fail", func() {
			err := indexedPriorityQueue.DeleteValue(-1)
			So(err, ShouldBeError, errors.New("index out of range"))
			err = indexedPriorityQueue.DeleteValue(11)
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("Delete By Index That Contains No Value Should Fail", func() {
			err := indexedPriorityQueue.DeleteValue(3)
			So(err, ShouldBeError, errors.New("index does not have value"))
		})
		Convey("Delete By Correct Index Should Work", func() {
			err := indexedPriorityQueue.DeleteValue(6)
			So(err, ShouldBeNil)
			So(indexedPriorityQueue.pq[3], ShouldEqual, 2)
			So(indexedPriorityQueue.pq[7], ShouldEqual, 7)
			So(indexedPriorityQueue.qp[6], ShouldEqual, -1)
		})
	})
}
