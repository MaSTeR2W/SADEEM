package absPath

import (
	"path/filepath"
	"runtime"
	"strings"
)

func get(stepBack int) string {
	var _, file, _, ok = runtime.Caller(stepBack)

	if !ok {
		panic("Can not Get Path")
	}
	return strings.ReplaceAll(filepath.Dir(file), "\\", "/")
}

func ToMe() string {
	return get(2)
}
