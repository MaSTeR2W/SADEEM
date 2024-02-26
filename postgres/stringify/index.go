package stringify

import (
	"reflect"
	"strconv"
	"strings"
)

func EscapeQuotes(val string) string {
	return strings.Replace(val, "'", "''", -1)
}

func Val(v any) string {
	switch v := v.(type) {
	case string:
		return "'" + EscapeQuotes(v) + "'"
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.FormatInt(int64(v), 10)
	case int16:
		return strconv.FormatInt(int64(v), 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	case uint8:
		return strconv.FormatUint(uint64(v), 10)
	case uint16:
		return strconv.FormatUint(uint64(v), 10)
	case uint32:
		return strconv.FormatUint(uint64(v), 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case bool:
		var b = "false"
		if v {
			b = "true"
		}
		return b

	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case []string:
		var vals = make([]string, 0, len(v))
		for _, e := range v {
			vals = append(vals, "'"+EscapeQuotes(e)+"'")
		}
		return "{" + strings.Join(vals, ",") + "}"
	case []int:
		return SliceOfInts[int](v)
	case []int8:
		return SliceOfInts[int8](v)
	case []int16:
		return SliceOfInts[int16](v)
	case []int32:
		return SliceOfInts[int32](v)
	case []int64:
		return SliceOfInts[int64](v)
	case []uint:
		return SliceOfUints[uint](v)
	case []uint8:
		return SliceOfUints[uint8](v)
	case []uint16:
		return SliceOfUints[uint16](v)
	case []uint32:
		return SliceOfUints[uint32](v)
	case []uint64:
		return SliceOfUints[uint64](v)
	default:
		panic("Can not handle the following type: " + reflect.TypeOf(v).Kind().String())
	}
}
