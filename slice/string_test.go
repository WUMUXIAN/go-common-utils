package slice

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestString(t *testing.T) {
	var index int
	var boolean bool
	x := []string{"1", "2", "3", "4", "5"}
	y := []string{"6", "7", "8", "9", "10"}
	z := []string{"6", "7", "8", "9", "10", "11"}

	Convey("Test The String Slice Functions\n", t, func() {
		Convey("IndexOfString", func() {
			index = IndexOfString(x, "1")
			So(index, ShouldEqual, 0)
			index = IndexOfString(x, "6")
			So(index, ShouldEqual, -1)
		})

		Convey("ContainsString", func() {
			boolean = ContainsString(x, "1")
			So(boolean, ShouldBeTrue)
			boolean = ContainsString(x, "6")
			So(boolean, ShouldBeFalse)
		})

		Convey("EqualsStrings", func() {
			boolean = EqualsStrings(x, x)
			So(boolean, ShouldBeTrue)
			boolean = EqualsStrings(x, y)
			So(boolean, ShouldBeFalse)
			boolean = EqualsStrings(x, z)
			So(boolean, ShouldBeFalse)
			boolean = EqualsStrings(y, z)
			So(boolean, ShouldBeFalse)
		})

		Convey("CopyStrings", func() {
			x1 := CopyStrings(x)
			boolean = EqualsStrings(x, x1)
			So(boolean, ShouldBeTrue)
		})

		Convey("CutStrings", func() {
			x1, err := CutStrings(x, 2, 4)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"1", "2", "5"})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
			_, err = CutStrings(x, -1, 6)
			So(err, ShouldNotBeNil)
			_, err = CutStrings(x, 6, 5)
			So(err, ShouldNotBeNil)
		})

		Convey("RemoveString", func() {
			x1 := RemoveString(x, "3")
			So(x1, ShouldResemble, []string{"1", "2", "4", "5"})
			x2 := RemoveString(x1, "10")
			So(x2, ShouldResemble, []string{"1", "2", "4", "5"})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("RemoveStringAt", func() {
			x1, err := RemoveStringAt(x, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"1", "2", "3", "5"})
			_, err = RemoveStringAt(x, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("InsertStringAt", func() {
			x1, err := InsertStringAt(x, "6", 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"1", "2", "3", "6", "4", "5"})
			_, err = InsertStringAt(x, "6", 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("InsertStringsAt", func() {
			x1, err := InsertStringsAt(x, y, 3)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"1", "2", "3", "6", "7", "8", "9", "10", "4", "5"})
			_, err = InsertStringsAt(x, y, 8)
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("PopFirstString", func() {
			v, x1, err := PopFirstString(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"2", "3", "4", "5"})
			So(v, ShouldResemble, "1")
			_, _, err = PopFirstString([]string{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("PopLastString", func() {
			v, x1, err := PopLastString(x)
			So(err, ShouldBeNil)
			So(x1, ShouldResemble, []string{"1", "2", "3", "4"})
			So(v, ShouldResemble, "5")
			_, _, err = PopLastString([]string{})
			So(err, ShouldNotBeNil)
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("FilterStrings", func() {
			x1 := FilterStrings(x, func(v string) bool {
				return v == "2" || v == "3"
			})
			So(x1, ShouldResemble, []string{"2", "3"})
			x1 = FilterStrings(x, func(v string) bool {
				return v == "10"
			})
			So(x1, ShouldResemble, []string{})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("ReverseStrings", func() {
			x1 := ReverseStrings(x)
			So(x1, ShouldResemble, []string{"5", "4", "3", "2", "1"})
			// Since it's a in-place operation, x will be changed, we need to re-init it.
			x = []string{"1", "2", "3", "4", "5"}
		})

		Convey("ShuffleStrings", func() {
			x1 := ShuffleStrings(x)
			So(x1, ShouldNotResemble, []string{"1", "2", "3", "4", "5"})
		})

		Convey("MergeStrings", func() {
			x1 := []string{"1", "2", "3"}
			x2 := []string{"2", "3", "4", "5"}
			merged := MergeStrings(x1, x2)
			So(merged, ShouldResemble, []string{"1", "2", "3", "4", "5"})
			merged = MergeStrings(x1, x2, "1", "2", "5")
			So(merged, ShouldResemble, []string{"3", "4"})
		})

		Convey("UniqueStrings", func() {
			x1 := []string{"1", "2", "2", "3", "4", "5", "5"}
			uniqueStrings := UniqueStrings(x1)
			So(uniqueStrings, ShouldResemble, []string{"1", "2", "3", "4", "5"})
		})

		Convey("TransformStrings", func() {
			target := []string{"1", "2", "3"}
			current := []string{}
			add, remove := TransformStrings(target, current)
			So(add, ShouldHaveLength, 3)
			So(add, ShouldContain, "1")
			So(add, ShouldContain, "2")
			So(add, ShouldContain, "3")
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []string{})

			target = []string{}
			current = []string{"1", "2", "3"}
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []string{})
			So(remove, ShouldHaveLength, 3)
			So(remove, ShouldContain, "1")
			So(remove, ShouldContain, "2")
			So(remove, ShouldContain, "3")

			target = []string{"3", "4", "5"}
			current = []string{"1", "2", "3"}
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 2)
			So(add, ShouldContain, "4")
			So(add, ShouldContain, "5")
			So(remove, ShouldHaveLength, 2)
			So(remove, ShouldContain, "1")
			So(remove, ShouldContain, "2")

			target = []string{}
			current = []string{}
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []string{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []string{})

			target = []string{"1", "2", "3"}
			current = []string{"1", "2", "3"}
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []string{})
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []string{})

			target = nil
			current = []string{"1146851694"}
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 0)
			So(add, ShouldResemble, []string{})
			So(remove, ShouldHaveLength, 1)
			So(remove, ShouldContain, "1146851694")

			target = []string{"1146851694"}
			current = nil
			add, remove = TransformStrings(target, current)
			So(add, ShouldHaveLength, 1)
			So(add, ShouldContain, "1146851694")
			So(remove, ShouldHaveLength, 0)
			So(remove, ShouldResemble, []string{})
		})

		Convey("ChunkStrings", func() {
			target := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks := ChunkStrings(target, 3)
			So(chunks, ShouldHaveLength, 4)
			So(chunks[0], ShouldResemble, []string{"1", "2", "3"})
			So(chunks[1], ShouldResemble, []string{"4", "5", "6"})
			So(chunks[2], ShouldResemble, []string{"7", "8", "9"})
			So(chunks[3], ShouldResemble, []string{"0"})

			target = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks = ChunkStrings(target, 4)
			So(chunks, ShouldHaveLength, 3)
			So(chunks[0], ShouldResemble, []string{"1", "2", "3", "4"})
			So(chunks[1], ShouldResemble, []string{"5", "6", "7", "8"})
			So(chunks[2], ShouldResemble, []string{"9", "0"})

			target = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks = ChunkStrings(target, 1)
			So(chunks, ShouldHaveLength, 10)
			So(chunks[0], ShouldResemble, []string{"1"})
			So(chunks[1], ShouldResemble, []string{"2"})
			So(chunks[2], ShouldResemble, []string{"3"})
			So(chunks[3], ShouldResemble, []string{"4"})
			So(chunks[4], ShouldResemble, []string{"5"})
			So(chunks[5], ShouldResemble, []string{"6"})
			So(chunks[6], ShouldResemble, []string{"7"})
			So(chunks[7], ShouldResemble, []string{"8"})
			So(chunks[8], ShouldResemble, []string{"9"})
			So(chunks[9], ShouldResemble, []string{"0"})

			target = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks = ChunkStrings(target, 10)
			So(chunks, ShouldHaveLength, 1)
			So(chunks[0], ShouldResemble, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"})

			target = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks = ChunkStrings(target, 9)
			So(chunks, ShouldHaveLength, 2)
			So(chunks[0], ShouldResemble, []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"})
			So(chunks[1], ShouldResemble, []string{"0"})

			target = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
			chunks = ChunkStrings(target, 0)
			So(chunks, ShouldHaveLength, 0)
		})
	})
}
