package pgHprs

import (
	"strconv"
	"strings"

	"github.com/MaSTeR2W/SADEEM/postgres/stringify"
)

func EscapeQuotes(val string) string {
	return strings.Replace(val, "'", "''", -1)
}

func UpdateSet(m map[string]any) string {
	var sets = make([]string, 0, len(m))
	for key, val := range m {
		sets = append(sets, key+"="+stringify.Val(val))
	}
	return strings.Join(sets, ",")
}

func BuildSqlUpdate(sql string, vals map[string]any) string {
	mLen := len(vals)
	var c int = 1
	for key, val := range vals {
		if strVal, ok := val.(string); ok {
			sql += key + "='" + EscapeQuotes(strVal) + "'"
		} else {
			val, _ := val.(int64)
			sql += key + "=" + strconv.FormatInt(val, 10)
		}
		if c < mLen {
			sql += ","
		}
	}
	return sql
}

func SetBindVars(keys ...string) string {
	var str = ""
	for i, key := range keys {
		str += "," + key + "=$" + strconv.FormatInt(int64(i+1), 10)
	}

	return str[1:]
}
