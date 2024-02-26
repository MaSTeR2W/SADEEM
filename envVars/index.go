package envVars

import (
	"github.com/joho/godotenv"
)

var DidRead = false

func init() {
	Read()
}

func Read() {
	if DidRead {
		return
	}
	godotenv.Load(".env")
	DidRead = true
}
