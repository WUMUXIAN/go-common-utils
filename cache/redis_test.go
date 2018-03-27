package cache

import (
	"sort"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type TestValues struct {
	A string
	B int
	C int64
}

func TestRedis(t *testing.T) {
	var err error

	err = NewRedisCacher("127.0.0.1:6379", "")
	Convey("New Redis Cacher Should Be OK", t, func() {
		So(err, ShouldBeNil)
	})

	GobRegister(&TestValues{})

	err = Redis.Set("testKey1", "testValue", 300)
	Convey("Set String Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.Set("testKey2", 1, 300)
	Convey("Set Int Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.Set("testKey3", 1.5, 300)
	Convey("Set Float Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.SetGob("testKey4", &TestValues{"A", 1, int64(1)}, 300)
	Convey("Set Composite Values By Gob Should Be OK", t, func() {
		So(err, ShouldBeNil)
	})

	testValue4, err := Redis.GetGob("testKey4")
	Convey("Get Composite Values By Gob Should Be OK", t, func() {
		So(err, ShouldBeNil)
		testValues, ok := testValue4.(*TestValues)
		So(ok, ShouldBeTrue)
		So(testValues.A, ShouldEqual, "A")
		So(testValues.B, ShouldEqual, 1)
		So(testValues.C, ShouldEqual, int64(1))
	})

	testValue3, err := Redis.GetFloat64("testKey3")
	Convey("Get Float Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
		So(testValue3, ShouldEqual, float64(1.5))
	})

	testValue2, err := Redis.GetInt64("testKey2")
	Convey("Get Int Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
		So(testValue2, ShouldEqual, int64(1))
	})

	testValue1, err := Redis.GetString("testKey1")
	Convey("Get Int Value Should Be OK", t, func() {
		So(err, ShouldBeNil)
		So(testValue1, ShouldEqual, "testValue")
	})

	nextCursor, keys, err := Redis.Scan(0, 100, "testKey*")
	Convey("Scanning The Keys Should Be OK", t, func() {
		So(nextCursor, ShouldEqual, 0)
		So(err, ShouldBeNil)
		sort.Strings(keys)
		So(keys, ShouldResemble, []string{"testKey1", "testKey2", "testKey3", "testKey4"})
	})

	Redis.Del("testKey1")
	Redis.Del("testKey2")
	Redis.Del("testKey3")
	Redis.Del("testKey4")

	nextCursor, keys, err = Redis.Scan(0, 100, "testKey*")
	Convey("Scanning The Keys Should Be OK", t, func() {
		So(nextCursor, ShouldEqual, 0)
		So(err, ShouldBeNil)
		So(keys, ShouldBeEmpty)
	})
}
