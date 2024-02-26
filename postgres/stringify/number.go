package stringify

import "strconv"

func Ints[T int | int8 | int16 | int32 | int64](num T) string {
	if v, ok := any(num).(int64); ok {
		return strconv.FormatInt(v, 10)
	} else {
		return strconv.FormatInt(int64(num), 10)

	}
}

func Uints[T uint | uint8 | uint16 | uint32 | uint64](num T) string {
	if v, ok := any(num).(uint64); ok {
		return strconv.FormatUint(v, 10)
	} else {
		return strconv.FormatUint(uint64(num), 10)

	}
}
