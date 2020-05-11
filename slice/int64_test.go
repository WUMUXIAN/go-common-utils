package slice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt64(t *testing.T) {
	var index int
	var boolean bool
	x := []int64{1, 2, 3, 4, 5}
	y := []int64{6, 7, 8, 9, 10}
	z := []int64{11, 12, 13, 14, 15, 16}

	Convey("Test The Int64 Slice Functions\n", t, func() {
		Convey("IndexOfInt64", func() {
			index = IndexOfInt64(x, 1)
			So(index, ShouldEqual, 0)
			index = IndexOfInt64(x, 6)
			So(index, ShouldEqual, -1)
		})

		Convey("ContainsInt64", func() {
			boolean = ContainsInt64(x, 1)
			So(boolean, ShouldBeTrue)
			boolean = ContainsInt64(x, 6)
			So(boolean, ShouldBeFalse)
		})

		Convey("EqualsInt64s", func() {
			boolean = EqualsInt64s(x, x)
			So(boolean, ShouldBeTrue)
			boolean = EqualsInt64s(x, y)
			So(boolean, ShouldBeFalse)
			boolean = EqualsInt64s(x, z)
			So(boolean, ShouldBeFalse)
			boolean = EqualsInt64s(y, z)
			So(boolean, ShouldBeFalse)
		})

		Convey("CopyInt64s", func() {
			x1 := CopyInt64s(x)
			boolean = EqualsInt64s(x, x1)
			So(boolean, ShouldBeTrue)
		})

		Convey("CutInt64s", func() {
			x1, err := CutInt64s(x, 2, 4)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{1, 2, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
			_, err = CutInt64s(x, -1, 6)
			So(err, ShouldNotBeNil)
			_, err = CutInt64s(x, 5, 4)
			So(err, ShouldNotBeNil)
		})

		Convey("RemoveInt64", func() {
			x1 := RemoveInt64(x, 3)
			So(x1, ShouldResemble, []int64{1, 2, 4, 5})
			x2 := RemoveInt64(x1, 10)
			So(x2, ShouldResemble, []int64{1, 2, 4, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("RemoveInt64At", func() {
			x1, err := RemoveInt64At(x, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{1, 2, 3, 5})
			_, err = RemoveInt64At(x, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("InsertInt64At", func() {
			x1, err := InsertInt64At(x, 6, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{1, 2, 3, 6, 4, 5})
			_, err = InsertInt64At(x, 6, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("InsertInt64sAt", func() {
			x1, err := InsertInt64sAt(x, y, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{1, 2, 3, 6, 7, 8, 9, 10, 4, 5})
			_, err = InsertInt64sAt(x, y, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("PopFirstInt64", func() {
			v, x1, err := PopFirstInt64(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{2, 3, 4, 5})
			So(v, ShouldResemble, int64(1))
			_, _, err = PopFirstInt64([]int64{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("PopLastInt64", func() {
			v, x1, err := PopLastInt64(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int64{1, 2, 3, 4})
			So(v, ShouldResemble, int64(5))
			_, _, err = PopLastInt64([]int64{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("FilterInt64s", func() {
			x1 := FilterInt64s(x, func(v int64) bool {
				return v == 2 || v == 3
			})
			So(x1, ShouldResemble, []int64{2, 3})
			x1 = FilterInt64s(x, func(v int64) bool {
				return v == 10
			})
			So(x1, ShouldResemble, []int64{})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("ReverseInt64s", func() {
			x1 := ReverseInt64s(x)
			So(x1, ShouldResemble, []int64{5, 4, 3, 2, 1})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int64{1, 2, 3, 4, 5}
		})

		Convey("ShuffleInt64s", func() {
			x1 := ShuffleInt64s(x)
			So(x1, ShouldNotResemble, []int64{1, 2, 3, 4, 5})
		})

		Convey("MergeInt64s", func() {
			x1 := []int64{1, 2, 3}
			x2 := []int64{2, 3, 4, 5}
			merged := MergeInt64s(x1, x2)
			So(merged, ShouldResemble, []int64{1, 2, 3, 4, 5})
			merged = MergeInt64s(x1, x2, 1, 2, 5)
			So(merged, ShouldResemble, []int64{3, 4})
		})

		Convey("SumOfInt64s", func() {
			x := []int64{1, 2, 3}
			sum := SumOfInt64s(x)
			So(sum, ShouldResemble, int64(6))
			x = []int64{}
			sum = SumOfInt64s(x)
			So(sum, ShouldResemble, int64(0))
			x = []int64{0, 0, 0}
			sum = SumOfInt64s(x)
			So(sum, ShouldResemble, int64(0))
		})

		Convey("TransformInt64s", func() {
			target := []int64{1, 2, 3}
			current := []int64{}
			add, remove := TransformInt64s(target, current)
			So(add, ShouldHaveLength, 3)
			So(add, ShouldContain, int64(1))
			So(add, ShouldContain, int64(2))
			So(add, ShouldContain, int64(3))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int64{})

			target = []int64{}
			current = []int64{1, 2, 3}
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int64{})
			So(remove, ShouldHaveLength, 3)
			So(remove, ShouldContain, int64(1))
			So(remove, ShouldContain, int64(2))
			So(remove, ShouldContain, int64(3))

			target = []int64{3, 4, 5}
			current = []int64{1, 2, 3}
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 2)
			So(add, ShouldContain, int64(4))
			So(add, ShouldContain, int64(5))
			So(remove, ShouldHaveLength, 2)
			So(remove, ShouldContain, int64(1))
			So(remove, ShouldContain, int64(2))

			target = []int64{}
			current = []int64{}
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int64{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int64{})

			target = []int64{1, 2, 3}
			current = []int64{1, 2, 3}
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int64{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int64{})

			target = nil
			current = []int64{1146851694}
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int64{})
			So(remove, ShouldHaveLength, 1)
			So(remove, ShouldContain, int64(1146851694))

			target = []int64{1146851694}
			current = nil
			add, remove = TransformInt64s(target, current)
			So(add, ShouldHaveLength, 1)
			So(add, ShouldContain, int64(1146851694))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int64{})
		})
	})
}
