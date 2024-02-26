package hprFns

import (
	"encoding/json"
	"fmt"
)

func IfEmpty(str1 string, str2 string) string {
	if str1 == "" {
		return str2
	}
	return str1
}

func PrintAll(v any) {
	js, err := json.MarshalIndent(v, "", "  ")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(js))
}
