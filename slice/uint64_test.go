package slice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUInt64(t *testing.T) {
	var index int
	var boolean bool
	x := []uint64{1, 2, 3, 4, 5}
	y := []uint64{6, 7, 8, 9, 10}
	z := []uint64{11, 12, 13, 14, 15, 16}

	Convey("Test The UInt64 Slice Functions\n", t, func() {
		Convey("IndexOfUInt64", func() {
			index = IndexOfUInt64(x, 1)
			So(index, ShouldEqual, 0)
			index = IndexOfUInt64(x, 6)
			So(index, ShouldEqual, -1)
		})

		Convey("ContainsUInt64", func() {
			boolean = ContainsUInt64(x, 1)
			So(boolean, ShouldBeTrue)
			boolean = ContainsUInt64(x, 6)
			So(boolean, ShouldBeFalse)
		})

		Convey("EqualsUInt64s", func() {
			boolean = EqualsUInt64s(x, x)
			So(boolean, ShouldBeTrue)
			boolean = EqualsUInt64s(x, y)
			So(boolean, ShouldBeFalse)
			boolean = EqualsUInt64s(x, z)
			So(boolean, ShouldBeFalse)
			boolean = EqualsUInt64s(y, z)
			So(boolean, ShouldBeFalse)
		})

		Convey("CopyUInt64s", func() {
			x1 := CopyUInt64s(x)
			boolean = EqualsUInt64s(x, x1)
			So(boolean, ShouldBeTrue)
		})

		Convey("CutUInt64s", func() {
			x1, err := CutUInt64s(x, 2, 4)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{1, 2, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
			_, err = CutUInt64s(x, -1, 6)
			So(err, ShouldNotBeNil)
			_, err = CutUInt64s(x, 5, 4)
			So(err, ShouldNotBeNil)
		})

		Convey("RemoveUInt64", func() {
			x1 := RemoveUInt64(x, 3)
			So(x1, ShouldResemble, []uint64{1, 2, 4, 5})
			x2 := RemoveUInt64(x1, 10)
			So(x2, ShouldResemble, []uint64{1, 2, 4, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("RemoveUInt64At", func() {
			x1, err := RemoveUInt64At(x, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{1, 2, 3, 5})
			_, err = RemoveUInt64At(x, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("InsertUInt64At", func() {
			x1, err := InsertUInt64At(x, 6, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{1, 2, 3, 6, 4, 5})
			_, err = InsertUInt64At(x, 6, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("InsertUInt64sAt", func() {
			x1, err := InsertUInt64sAt(x, y, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{1, 2, 3, 6, 7, 8, 9, 10, 4, 5})
			_, err = InsertUInt64sAt(x, y, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("PopFirstUInt64", func() {
			v, x1, err := PopFirstUInt64(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{2, 3, 4, 5})
			So(v, ShouldResemble, uint64(1))
			_, _, err = PopFirstUInt64([]uint64{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("PopLastUInt64", func() {
			v, x1, err := PopLastUInt64(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []uint64{1, 2, 3, 4})
			So(v, ShouldResemble, uint64(5))
			_, _, err = PopLastUInt64([]uint64{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("FilterUInt64s", func() {
			x1 := FilterUInt64s(x, func(v uint64) bool {
				return v == 2 || v == 3
			})
			So(x1, ShouldResemble, []uint64{2, 3})
			x1 = FilterUInt64s(x, func(v uint64) bool {
				return v == 10
			})
			So(x1, ShouldResemble, []uint64{})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("ReverseUInt64s", func() {
			x1 := ReverseUInt64s(x)
			So(x1, ShouldResemble, []uint64{5, 4, 3, 2, 1})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []uint64{1, 2, 3, 4, 5}
		})

		Convey("ShuffleUInt64s", func() {
			x1 := ShuffleUInt64s(x)
			So(x1, ShouldNotResemble, []uint64{1, 2, 3, 4, 5})
		})

		Convey("MergeUInt64s", func() {
			x1 := []uint64{1, 2, 3}
			x2 := []uint64{2, 3, 4, 5}
			merged := MergeUInt64s(x1, x2)
			So(merged, ShouldResemble, []uint64{1, 2, 3, 4, 5})
			merged = MergeUInt64s(x1, x2, 1, 2, 5)
			So(merged, ShouldResemble, []uint64{3, 4})
		})

		Convey("UniqueUInt64s", func() {
			x1 := []uint64{1, 2, 2, 3, 4, 5, 5}
			uniqueUInt64s := UniqueUInt64s(x1)
			So(uniqueUInt64s, ShouldResemble, []uint64{1, 2, 3, 4, 5})
		})

		Convey("SumOfUInt64s", func() {
			x := []uint64{1, 2, 3}
			sum := SumOfUInt64s(x)
			So(sum, ShouldResemble, uint64(6))
			x = []uint64{}
			sum = SumOfUInt64s(x)
			So(sum, ShouldResemble, uint64(0))
			x = []uint64{0, 0, 0}
			sum = SumOfUInt64s(x)
			So(sum, ShouldResemble, uint64(0))
		})

		Convey("TransformUInt64s", func() {
			target := []uint64{1, 2, 3}
			current := []uint64{}
			add, remove := TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 3)
			So(add, ShouldContain, uint64(1))
			So(add, ShouldContain, uint64(2))
			So(add, ShouldContain, uint64(3))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []uint64{})

			target = []uint64{}
			current = []uint64{1, 2, 3}
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []uint64{})
			So(remove, ShouldHaveLength, 3)
			So(remove, ShouldContain, uint64(1))
			So(remove, ShouldContain, uint64(2))
			So(remove, ShouldContain, uint64(3))

			target = []uint64{3, 4, 5}
			current = []uint64{1, 2, 3}
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 2)
			So(add, ShouldContain, uint64(4))
			So(add, ShouldContain, uint64(5))
			So(remove, ShouldHaveLength, 2)
			So(remove, ShouldContain, uint64(1))
			So(remove, ShouldContain, uint64(2))

			target = []uint64{}
			current = []uint64{}
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []uint64{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []uint64{})

			target = []uint64{1, 2, 3}
			current = []uint64{1, 2, 3}
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []uint64{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []uint64{})

			target = nil
			current = []uint64{1146851694}
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []uint64{})
			So(remove, ShouldHaveLength, 1)
			So(remove, ShouldContain, uint64(1146851694))

			target = []uint64{1146851694}
			current = nil
			add, remove = TransformUInt64s(target, current)
			So(add, ShouldHaveLength, 1)
			So(add, ShouldContain, uint64(1146851694))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []uint64{})
		})
	})
}
