package cache

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"sort"
	"testing"

	"github.com/gomodule/redigo/redis"
	. "github.com/smartystreets/goconvey/convey"
)

type TestValues struct {
	A string
	B int
	C int64
}

func TestRedis(t *testing.T) {
	var err error

	// Point to wring server address or port
	err = NewRedisCacher("127.0.0.1:6380", "")
	Convey("New Redis Cacher Should Be Failed\n", t, func() {
		So(err, ShouldNotBeNil)
	})

	err = NewRedisCacher("127.0.0.2:6379", "")
	Convey("New Redis Cacher Should Be Failed\n", t, func() {
		So(err, ShouldNotBeNil)
	})

	// Wrong password
	err = NewRedisCacher("127.0.0.1:6379", "WrongPassword")
	Convey("New Redis Cacher Should Be Failed\n", t, func() {
		So(err, ShouldNotBeNil)
	})

	err = NewRedisCacher("127.0.0.1:6379", "", 1)
	Convey("New Redis Cacher Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	// Let's new another cacher immediately, it should be ok.
	err = NewRedisCacher("127.0.0.1:6379", "", 1)
	Convey("Double New Redis Cacher Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	GobRegister(&TestValues{})

	err = Redis.Set("testKey1", "testValue", 300)
	Convey("Set String Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.Set("testKey2", 1, 300)
	Convey("Set Int Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.Set("testKey3", 1.5, 300)
	Convey("Set Float Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.SetGob("testKey4", &TestValues{"A", 1, int64(1)}, 300)
	Convey("Set Composite Values By Gob Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.SetJSON("testKey5", &TestValues{"A", 1, int64(1)}, 300)
	Convey("Set Composite Values By JSON Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.Expire("testKey5", 500)
	Convey("Update Expire For Key Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	ttl, err := Redis.TTL("testKey5")
	Convey("Get Expire For Key Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(ttl, ShouldBeBetween, 498, 501)
	})

	testValue5, err := Redis.GetJSON("testKey5")
	Convey("Get Composite Values By JSON Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		var testValues TestValues
		err = json.Unmarshal(testValue5, &testValues)
		So(err, ShouldBeNil)
		So(testValues.A, ShouldEqual, "A")
		So(testValues.B, ShouldEqual, 1)
		So(testValues.C, ShouldEqual, int64(1))
	})

	testValue4, err := Redis.GetGob("testKey4")
	Convey("Get Composite Values By Gob Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		testValues, ok := testValue4.(*TestValues)
		So(ok, ShouldBeTrue)
		So(testValues.A, ShouldEqual, "A")
		So(testValues.B, ShouldEqual, 1)
		So(testValues.C, ShouldEqual, int64(1))
	})

	testValue3, err := Redis.GetFloat64("testKey3")
	Convey("Get Float Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(testValue3, ShouldEqual, float64(1.5))
	})

	testValue2, err := Redis.GetInt64("testKey2")
	Convey("Get Int Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(testValue2, ShouldEqual, int64(1))
	})

	testValue1, err := Redis.GetString("testKey1")
	Convey("Get Int Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(testValue1, ShouldEqual, "testValue")
	})

	values, err := Redis.MultipleGet("testKey1", "testKey2", "testKey3", "testKey4", "testKey5")
	Convey("Get Values In Bulk Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		var value1 string
		var value2 int64
		var value3 float64
		values, err = redis.Scan(values, &value1, &value2, &value3)
		So(value1, ShouldEqual, "testValue")
		So(value2, ShouldEqual, int64(1))
		So(value3, ShouldEqual, float64(1.5))

		b, err := redis.Bytes(values[0], nil)
		value := make(map[interface{}]interface{})
		if err == nil {
			buffer := bytes.NewBuffer(b)
			dec := gob.NewDecoder(buffer)
			err = dec.Decode(&value)
		}
		testValues, ok := value["value"].(*TestValues)
		So(ok, ShouldBeTrue)
		So(testValues.A, ShouldEqual, "A")
		So(testValues.B, ShouldEqual, 1)
		So(testValues.C, ShouldEqual, int64(1))

		b, err = redis.Bytes(values[1], nil)
		if err == nil {
			var testValues TestValues
			err = json.Unmarshal(testValue5, &testValues)
			So(err, ShouldBeNil)
			So(testValues.A, ShouldEqual, "A")
			So(testValues.B, ShouldEqual, 1)
			So(testValues.C, ShouldEqual, int64(1))
		}
	})

	jsonBytes, err := Redis.MultipleGetJSON("testKey5")
	Convey("Get JSON Values In Bulk Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		var testValues TestValues
		So(len(jsonBytes), ShouldEqual, 1)
		err = json.Unmarshal(jsonBytes[0], &testValues)
		So(err, ShouldBeNil)
		So(testValues.A, ShouldEqual, "A")
		So(testValues.B, ShouldEqual, 1)
		So(testValues.C, ShouldEqual, int64(1))

	})

	nextCursor, keys, err := Redis.Scan(0, 100, "testKey*")
	Convey("Scanning The Keys Should Be OK\n", t, func() {
		So(nextCursor, ShouldBeGreaterThanOrEqualTo, 0)
		So(err, ShouldBeNil)
		sort.Strings(keys)
		So(keys, ShouldResemble, []string{"testKey1", "testKey2", "testKey3", "testKey4", "testKey5"})
	})

	// Error case.
	nextCursor, keys, err = Redis.Scan(0, -100, "testKey*")
	Convey("Scanning The Keys With Wrong Count Should Not Be OK\n", t, func() {
		So(nextCursor, ShouldBeGreaterThanOrEqualTo, 0)
		So(err, ShouldNotBeNil)
	})

	Redis.Del("testKey1")
	Redis.Del("testKey2")
	Redis.Del("testKey3")
	Redis.Del("testKey4")
	Redis.Del("testKey5")

	nextCursor, keys, err = Redis.Scan(0, 100, "testKey*")
	Convey("Scanning The Keys Should Be OK\n", t, func() {
		So(nextCursor, ShouldBeGreaterThanOrEqualTo, 0)
		So(err, ShouldBeNil)
		So(keys, ShouldBeEmpty)
	})

	err = Redis.HSet("hash1", "key1", 1)
	Convey("Set Hash Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	err = Redis.HSet("hash2", "key2", 1, 10)
	Convey("Set Hash Set With Expires Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err := redis.Int64(Redis.HGet("hash1", "key1"))
	Convey("Get Hash Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(1))
	})

	err = Redis.HINCRBY("hash1", "key1", 2)
	Convey("Increment Hash Set Value By Key Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err = redis.Int64(Redis.HGet("hash1", "key1"))
	Convey("Get Hash Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(3))
	})

	_ = Redis.HSet("hash-m", "key-m-1", 1)
	_ = Redis.HSet("hash-m", "key-m-2", 2)
	_ = Redis.HSet("hash-m", "key-m-3", 3)
	multiValues, err := redis.Int64s(Redis.HMGet("hash-m", "key-m-1", "key-m-2", "key-m-3"))
	Convey("Get Multiple Values For Hash Keys Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(multiValues, ShouldHaveLength, 3)
		So(multiValues[0], ShouldEqual, int64(1))
		So(multiValues[1], ShouldEqual, int64(2))
		So(multiValues[2], ShouldEqual, int64(3))
	})

	multiValues, err = redis.Int64s(Redis.HMGet("hash-m", "key-m-2", "key-m-3", "key-m-1"))
	Convey("Get Multiple Values Order For Hash Keys Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(multiValues, ShouldHaveLength, 3)
		So(multiValues[0], ShouldEqual, int64(2))
		So(multiValues[1], ShouldEqual, int64(3))
		So(multiValues[2], ShouldEqual, int64(1))
	})

	err = Redis.Flush()
	Convey("Flush All Keys In Current DB Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err = redis.Int64(Redis.GetDBSize())
	Convey("Get Size After Flush Should Return 0\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(0))
	})

	err = Redis.Set("testKey1", "testValue", 300)
	Convey("Set String Value Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err = redis.Int64(Redis.GetDBSize())
	Convey("Get Size After Set A Key Should Return 1\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(1))
	})

	err = Redis.HSet("hash1", "key1", 1)
	Convey("Set Hash Set Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err = redis.Int64(Redis.HDel("hash1", "key1"))
	Convey("Delete Hash Key Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(1))
	})

	value, err = redis.Int64(Redis.HGet("hash1", "key1"))
	Convey("Get Hash Key Should Return Nil\n", t, func() {
		So(err, ShouldBeError, errors.New("redigo: nil returned"))
		So(value, ShouldEqual, int64(0))
	})

	err = Redis.FlushAll()
	Convey("Flush All Keys Should Be OK\n", t, func() {
		So(err, ShouldBeNil)
	})

	value, err = redis.Int64(Redis.GetDBSize())
	Convey("Get Size After Flush Should Return 0\n", t, func() {
		So(err, ShouldBeNil)
		So(value, ShouldEqual, int64(0))
	})
}
