package slice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestInt(t *testing.T) {
	var index int
	var boolean bool
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	z := []int{11, 12, 13, 14, 15, 16}

	Convey("Test The Int Slice Functions\n", t, func() {
		Convey("IndexOfInt", func() {
			index = IndexOfInt(x, 1)
			So(index, ShouldEqual, 0)
			index = IndexOfInt(x, 6)
			So(index, ShouldEqual, -1)
		})

		Convey("ContainsInt", func() {
			boolean = ContainsInt(x, 1)
			So(boolean, ShouldBeTrue)
			boolean = ContainsInt(x, 6)
			So(boolean, ShouldBeFalse)
		})

		Convey("EqualsInts", func() {
			boolean = EqualsInts(x, x)
			So(boolean, ShouldBeTrue)
			boolean = EqualsInts(x, y)
			So(boolean, ShouldBeFalse)
			boolean = EqualsInts(x, z)
			So(boolean, ShouldBeFalse)
			boolean = EqualsInts(y, z)
			So(boolean, ShouldBeFalse)
		})

		Convey("CopyInts", func() {
			x1 := CopyInts(x)
			boolean = EqualsInts(x, x1)
			So(boolean, ShouldBeTrue)
		})

		Convey("CutInts", func() {
			x1, err := CutInts(x, 2, 4)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{1, 2, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
			_, err = CutInts(x, -1, 6)
			So(err, ShouldNotBeNil)
			_, err = CutInts(x, 5, 4)
			So(err, ShouldNotBeNil)
		})

		Convey("RemoveInt", func() {
			x1 := RemoveInt(x, 3)
			So(x1, ShouldResemble, []int{1, 2, 4, 5})
			x2 := RemoveInt(x1, 10)
			So(x2, ShouldResemble, []int{1, 2, 4, 5})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("RemoveIntAt", func() {
			x1, err := RemoveIntAt(x, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{1, 2, 3, 5})
			_, err = RemoveIntAt(x, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("InsertIntAt", func() {
			x1, err := InsertIntAt(x, 6, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{1, 2, 3, 6, 4, 5})
			_, err = InsertIntAt(x, 6, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("InsertIntsAt", func() {
			x1, err := InsertIntsAt(x, y, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{1, 2, 3, 6, 7, 8, 9, 10, 4, 5})
			_, err = InsertIntsAt(x, y, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("PopFirstInt", func() {
			v, x1, err := PopFirstInt(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{2, 3, 4, 5})
			So(v, ShouldResemble, 1)
			_, _, err = PopFirstInt([]int{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("PopLastInt", func() {
			v, x1, err := PopLastInt(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []int{1, 2, 3, 4})
			So(v, ShouldResemble, 5)
			_, _, err = PopLastInt([]int{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("FilterInts", func() {
			x1 := FilterInts(x, func(v int) bool {
				return v == 2 || v == 3
			})
			So(x1, ShouldResemble, []int{2, 3})
			x1 = FilterInts(x, func(v int) bool {
				return v == 10
			})
			So(x1, ShouldResemble, []int{})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("ReverseInts", func() {
			x1 := ReverseInts(x)
			So(x1, ShouldResemble, []int{5, 4, 3, 2, 1})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []int{1, 2, 3, 4, 5}
		})

		Convey("ShuffleInts", func() {
			x1 := ShuffleInts(x)
			So(x1, ShouldNotResemble, []int{1, 2, 3, 4, 5})
		})

		Convey("MergeInts", func() {
			x1 := []int{1, 2, 3}
			x2 := []int{2, 3, 4, 5}
			merged := MergeInts(x1, x2)
			So(merged, ShouldResemble, []int{1, 2, 3, 4, 5})
			merged = MergeInts(x1, x2, 1, 2, 5)
			So(merged, ShouldResemble, []int{3, 4})
		})

		Convey("UniqueInts", func() {
			x1 := []int{1, 2, 2, 3, 4, 5, 5}
			uniqueInts := UniqueInts(x1)
			So(uniqueInts, ShouldResemble, []int{1, 2, 3, 4, 5})
		})

		Convey("SumOfInts", func() {
			x := []int{1, 2, 3}
			sum := SumOfInts(x)
			So(sum, ShouldResemble, 6)
			x = []int{}
			sum = SumOfInts(x)
			So(sum, ShouldResemble, 0)
			x = []int{0, 0, 0}
			sum = SumOfInts(x)
			So(sum, ShouldResemble, 0)
		})

		Convey("TransformInts", func() {
			target := []int{1, 2, 3}
			current := []int{}
			add, remove := TransformInts(target, current)
			So(add, ShouldHaveLength, 3)
			So(add, ShouldContain, int(1))
			So(add, ShouldContain, int(2))
			So(add, ShouldContain, int(3))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int{})

			target = []int{}
			current = []int{1, 2, 3}
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int{})
			So(remove, ShouldHaveLength, 3)
			So(remove, ShouldContain, int(1))
			So(remove, ShouldContain, int(2))
			So(remove, ShouldContain, int(3))

			target = []int{3, 4, 5}
			current = []int{1, 2, 3}
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 2)
			So(add, ShouldContain, int(4))
			So(add, ShouldContain, int(5))
			So(remove, ShouldHaveLength, 2)
			So(remove, ShouldContain, int(1))
			So(remove, ShouldContain, int(2))

			target = []int{}
			current = []int{}
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int{})

			target = []int{1, 2, 3}
			current = []int{1, 2, 3}
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int{})

			target = nil
			current = []int{1146851694}
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []int{})
			So(remove, ShouldHaveLength, 1)
			So(remove, ShouldContain, int(1146851694))

			target = []int{1146851694}
			current = nil
			add, remove = TransformInts(target, current)
			So(add, ShouldHaveLength, 1)
			So(add, ShouldContain, int(1146851694))
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []int{})
		})

		Convey("ChunkInts", func() {
			target := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks := ChunkInts(target, 3)
			So(chunks, ShouldHaveLength, 4)
			So(chunks[0], ShouldResemble, []int{1, 2, 3})
			So(chunks[1], ShouldResemble, []int{4, 5, 6})
			So(chunks[2], ShouldResemble, []int{7, 8, 9})
			So(chunks[3], ShouldResemble, []int{0})

			target = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks = ChunkInts(target, 4)
			So(chunks, ShouldHaveLength, 3)
			So(chunks[0], ShouldResemble, []int{1, 2, 3, 4})
			So(chunks[1], ShouldResemble, []int{5, 6, 7, 8})
			So(chunks[2], ShouldResemble, []int{9, 0})

			target = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks = ChunkInts(target, 1)
			So(chunks, ShouldHaveLength, 10)
			So(chunks[0], ShouldResemble, []int{1})
			So(chunks[1], ShouldResemble, []int{2})
			So(chunks[2], ShouldResemble, []int{3})
			So(chunks[3], ShouldResemble, []int{4})
			So(chunks[4], ShouldResemble, []int{5})
			So(chunks[5], ShouldResemble, []int{6})
			So(chunks[6], ShouldResemble, []int{7})
			So(chunks[7], ShouldResemble, []int{8})
			So(chunks[8], ShouldResemble, []int{9})
			So(chunks[9], ShouldResemble, []int{0})

			target = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks = ChunkInts(target, 10)
			So(chunks, ShouldHaveLength, 1)
			So(chunks[0], ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0})

			target = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks = ChunkInts(target, 9)
			So(chunks, ShouldHaveLength, 2)
			So(chunks[0], ShouldResemble, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
			So(chunks[1], ShouldResemble, []int{0})

			target = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
			chunks = ChunkInts(target, 0)
			So(chunks, ShouldHaveLength, 0)

			target = []int{}
			chunks = ChunkInts(target, 1)
			So(chunks, ShouldHaveLength, 0)

			target = nil
			chunks = ChunkInts(target, 1)
			So(chunks, ShouldHaveLength, 0)
		})
	})
}
