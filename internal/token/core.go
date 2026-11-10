//ff:func feature=token type=reader control=sequence
//ff:what core — a field without the punctuation and symbols around it: punctuation is a fixed boundary
package token

import (
	"strings"
	"unicode"
)

func core(field string) []rune {
	return []rune(strings.TrimFunc(field, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}))
}
