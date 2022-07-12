package param_demo

import (
	"strings"
	"unicode"
)

func stripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
