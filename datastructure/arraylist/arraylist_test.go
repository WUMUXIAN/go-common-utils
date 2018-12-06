package arraylist

import (
	"errors"
	"testing"

	"github.com/TectusDreamlab/go-common-utils/datastructure/shared"
	. "github.com/smartystreets/goconvey/convey"
)

func TestArrayList(t *testing.T) {
	var err error
	var index int
	var element interface{}

	arrayList := New(shared.IntComparator)
	Convey("Create A New ArrayList", t, func() {
		Convey("Should Has Empty Size", func() {
			So(arrayList.GetSize(), ShouldEqual, 0)
			So(arrayList.ToSlice(), ShouldHaveLength, 0)
		})
	})

	arrayList.ensureCapacity(5)
	Convey("Ensure Capacity With A Smaller Value", t, func() {
		Convey("Should Not Change Anything", func() {
			So(arrayList.capacity, ShouldEqual, 10)
		})
	})

	arrayList.Append(2)
	Convey("Append An Element", t, func() {
		Convey("Array Size Should Be 1", func() {
			So(arrayList.GetSize(), ShouldEqual, 1)
		})
		Convey("Should Have Elements As Expected", func() {
			So(arrayList.ToSlice(), ShouldResemble, []interface{}{2})
		})
	})

	err = arrayList.Add(0, 3)
	Convey("Add An Element", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Array Size Should Be 2", func() {
			So(arrayList.GetSize(), ShouldEqual, 2)
		})
		Convey("Should Have Elements As Expected", func() {
			So(arrayList.ToSlice(), ShouldResemble, []interface{}{3, 2})
		})
	})

	err = arrayList.Add(3, 3)
	Convey("Add An Element At A Wrong Index", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("Array Size Should Be 2", func() {
			So(arrayList.GetSize(), ShouldEqual, 2)
		})
		Convey("Should Have Elements As Expected", func() {
			So(arrayList.ToSlice(), ShouldResemble, []interface{}{3, 2})
		})
	})

	for i := 0; i < 100; i++ {
		arrayList.Append(i)
	}
	Convey("Add 100 Elements", t, func() {
		Convey("Array Size Should Be 102", func() {
			So(arrayList.GetSize(), ShouldEqual, 102)
		})
		Convey("Array Capcity Should Be 175", func() {
			So(arrayList.capacity, ShouldEqual, 175)
		})
	})

	err = arrayList.Remove(99)
	Convey("Remove An Element", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Array Size Should Be 101", func() {
			So(arrayList.GetSize(), ShouldEqual, 101)
		})
	})

	err = arrayList.Remove(10000)
	Convey("Remove An Non-existed Element", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("element not found"))
		})
		Convey("Array Size Should Be 101", func() {
			So(arrayList.GetSize(), ShouldEqual, 101)
		})
	})

	err = arrayList.RemoveAt(101)
	Convey("Remove At An Index Out Of Range", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("Array Size Should Be 101", func() {
			So(arrayList.GetSize(), ShouldEqual, 101)
		})
	})

	err = arrayList.RemoveAt(50)
	Convey("Remove An Element At Given Index", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Array Size Should Be 100", func() {
			So(arrayList.GetSize(), ShouldEqual, 100)
		})
		Convey("The New Element At Index 50 Should Be 100", func() {
			element, err := arrayList.Get(50)
			So(err, ShouldBeNil)
			So(element, ShouldEqual, 49)
		})
		Convey("The Element 48 Should Not Be In", func() {
			So(arrayList.Contains(48), ShouldBeFalse)
		})
	})

	index, err = arrayList.GetIndexOf(999)
	Convey("Get Index Of Non-Existed Element", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("element not found"))
		})
		Convey("Index Should Be -1", func() {
			So(index, ShouldEqual, -1)
		})
	})

	index, err = arrayList.GetIndexOf(32)
	Convey("Get Index Of An Element", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Index Should Be 34", func() {
			So(index, ShouldEqual, 34)
		})
	})

	element, err = arrayList.Get(200)
	Convey("Get Element At Out Of Range Index", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("index out of range"))
		})
		Convey("element Should Be Nil", func() {
			So(element, ShouldBeNil)
		})
	})

	element, err = arrayList.Get(7)
	Convey("Get Element At Given Index", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
		})
		Convey("Element Should Be 5", func() {
			So(element, ShouldEqual, 5)
		})
	})

	err = arrayList.Set(12321, 1233)
	Convey("Set Element At Out Of Range Index", t, func() {
		Convey("Should Not Be Successful", func() {
			So(err, ShouldBeError, errors.New("index out of range"))
		})
	})

	err = arrayList.Set(25, 1555)
	Convey("Set Element At Given Index", t, func() {
		Convey("Should Be Successful", func() {
			So(err, ShouldBeNil)
			element, err = arrayList.Get(25)
			So(err, ShouldBeNil)
			So(element, ShouldEqual, 1555)
		})
	})

	Convey("ArrayList Should Contain Element 27", t, func() {
		So(arrayList.Contains(27), ShouldBeTrue)
	})

	Convey("ArrayList Should Not Contain Element 1432", t, func() {
		So(arrayList.Contains(1432), ShouldBeFalse)
	})

	it := arrayList.Iterator()
	Convey("Test The Array List Iterator", t, func() {
		Convey("Get The First Element By Moving Next Once And Call Value", func() {
			So(it.Next(), ShouldBeTrue)
			So(it.Value(), ShouldEqual, 3)
		})
		Convey("Move Next 5 Times Should Work As Expected", func() {
			for i := 0; i < 5; i++ {
				So(it.Next(), ShouldBeTrue)
			}
			So(it.Value(), ShouldEqual, 3)
		})
		Convey("Move To The End Should Work As Expected", func() {
			it.End()
			So(it.Next(), ShouldBeFalse)
			So(it.Value(), ShouldBeNil)
		})
		Convey("Get The Last Element By Moving Prev Once And Call Value", func() {
			So(it.Prev(), ShouldBeTrue)
			So(it.Value(), ShouldEqual, 98)
		})
		Convey("Move Prev 5 Times Should Work As Expected", func() {
			for i := 0; i < 5; i++ {
				So(it.Prev(), ShouldBeTrue)
			}
			So(it.Value(), ShouldEqual, 93)
		})
		Convey("Move To The Begin Should Work As Expected", func() {
			it.Begin()
			So(it.Prev(), ShouldBeFalse)
			So(it.Value(), ShouldBeNil)
		})
	})

	arrayList.Clear()
	Convey("Clear The Array List", t, func() {
		Convey("Size Of Array List Should Be 0", func() {
			So(arrayList.GetSize(), ShouldEqual, 0)
		})
	})
}
