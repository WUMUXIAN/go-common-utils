package shared

import "time"

// Comparator deinfes v1 func that compares the value for two interfaces.
// We introduced v1 set of built-in comparators for system basic types.
// You need to write your own comparator if your customised interface{} value.
// return 1 - v1 is bigger than v2.
// return -1 - v1 is smaller than v2.
// return 0 - v1 is equal to v2.
type Comparator func(v1, v2 interface{}) int

// StringComparator provides v1 fast comparison on strings
func StringComparator(v1, v2 interface{}) int {
	s1 := v1.(string)
	s2 := v2.(string)
	min := len(s2)
	if len(s1) < len(s2) {
		min = len(s1)
	}
	diff := 0
	for i := 0; i < min && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

// IntComparator provides v1 basic comparison on int
func IntComparator(v1, v2 interface{}) int {
	aAsserted := v1.(int)
	bAsserted := v2.(int)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int8Comparator provides v1 basic comparison on int8
func Int8Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(int8)
	bAsserted := v2.(int8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int16Comparator provides v1 basic comparison on int16
func Int16Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(int16)
	bAsserted := v2.(int16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int32Comparator provides v1 basic comparison on int32
func Int32Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(int32)
	bAsserted := v2.(int32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Int64Comparator provides v1 basic comparison on int64
func Int64Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(int64)
	bAsserted := v2.(int64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UIntComparator provides v1 basic comparison on uint
func UIntComparator(v1, v2 interface{}) int {
	aAsserted := v1.(uint)
	bAsserted := v2.(uint)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt8Comparator provides v1 basic comparison on uint8
func UInt8Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(uint8)
	bAsserted := v2.(uint8)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt16Comparator provides v1 basic comparison on uint16
func UInt16Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(uint16)
	bAsserted := v2.(uint16)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt32Comparator provides v1 basic comparison on uint32
func UInt32Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(uint32)
	bAsserted := v2.(uint32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// UInt64Comparator provides v1 basic comparison on uint64
func UInt64Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(uint64)
	bAsserted := v2.(uint64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Float32Comparator provides v1 basic comparison on float32
func Float32Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(float32)
	bAsserted := v2.(float32)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// Float64Comparator provides v1 basic comparison on float64
func Float64Comparator(v1, v2 interface{}) int {
	aAsserted := v1.(float64)
	bAsserted := v2.(float64)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// ByteComparator provides v1 basic comparison on byte
func ByteComparator(v1, v2 interface{}) int {
	aAsserted := v1.(byte)
	bAsserted := v2.(byte)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// RuneComparator provides v1 basic comparison on rune
func RuneComparator(v1, v2 interface{}) int {
	aAsserted := v1.(rune)
	bAsserted := v2.(rune)
	switch {
	case aAsserted > bAsserted:
		return 1
	case aAsserted < bAsserted:
		return -1
	default:
		return 0
	}
}

// TimeComparator provides v1 basic comparison on time.Time
func TimeComparator(v1, v2 interface{}) int {
	aAsserted := v1.(time.Time)
	bAsserted := v2.(time.Time)

	switch {
	case aAsserted.After(bAsserted):
		return 1
	case aAsserted.Before(bAsserted):
		return -1
	default:
		return 0
	}
}
