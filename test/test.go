package main

import (
	"fmt"
	"regexp"
)

var allowedLocalPartChars = regexp.MustCompile(`[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+`)

func main() {

	fmt.Println(allowedLocalPartChars.MatchString("~"))
	fmt.Println(allowedLocalPartChars.MatchString("-"))
	fmt.Println(allowedLocalPartChars.MatchString("]"))
	fmt.Println(allowedLocalPartChars.MatchString("+"))
	fmt.Println(allowedLocalPartChars.MatchString("["))
	fmt.Println(allowedLocalPartChars.MatchString(";"))
	fmt.Println(allowedLocalPartChars.MatchString("َ"))
	fmt.Println(allowedLocalPartChars.MatchString("|"))
	fmt.Println(allowedLocalPartChars.MatchString("}"))
}
