package hprFns

import "strconv"

func Join[T uint | int](s []T, sep string) string {
	var str = ""
	switch ts := any(s).(type) {
	case []uint:
		var lIndex = len(ts) - 1
		for i := 0; i < lIndex; i++ {
			str += strconv.FormatUint(uint64(ts[i]), 10) + sep
		}

		return str + strconv.FormatUint(uint64(ts[lIndex]), 10)
	case []int:
		{
			var lIndex = len(ts) - 1
			for i := 0; i < lIndex; i++ {
				str += strconv.FormatInt(int64(ts[i]), 10) + sep
			}

			return str + strconv.FormatInt(int64(ts[lIndex]), 10)
		}
	}
	return str
}

// JoinPs: join with prefix and suffix
func JoinPs[T uint | int](s []T, prefix, suffix, sep string) string {
	var str = ""
	switch ts := any(s).(type) {
	case []uint:
		var lIndex = len(ts) - 1
		for i := 0; i < lIndex; i++ {
			str += prefix + strconv.FormatUint(uint64(ts[i]), 10) + suffix + sep
		}

		return str + prefix + strconv.FormatUint(uint64(ts[lIndex]), 10) + suffix
	case []int:
		{
			var lIndex = len(ts) - 1
			for i := 0; i < lIndex; i++ {
				str += prefix + strconv.FormatInt(int64(ts[i]), 10) + suffix + sep
			}

			return str + prefix + strconv.FormatInt(int64(ts[lIndex]), 10) + suffix
		}
	}
	return str
}
