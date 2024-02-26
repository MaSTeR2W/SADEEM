package stringify

import "strings"

func SliceOfInts[T int | int8 | int16 | int32 | int64](si []T) string {
	var str = make([]string, 0, len(si))

	for _, e := range si {
		str = append(str, Ints[T](e))
	}
	return "{" + strings.Join(str, ",") + "}"
}

func SliceOfUints[T uint | uint8 | uint16 | uint32 | uint64](si []T) string {
	var str = make([]string, 0, len(si))

	for _, e := range si {
		str = append(str, Uints[T](e))
	}
	return "{" + strings.Join(str, ",") + "}"
}
