//ff:func feature=token type=helper control=sequence
//ff:what trimSigns — a field without the punctuation marks, symbols and named signs (signNames) at its two ends, as letters; a named sign that is a letter for Unicode (ㆍ) is taken off too
package token

import (
	"strings"
	"unicode"
)

func trimSigns(field string) []rune {
	return []rune(strings.TrimFunc(field, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r) || len(signNames[r]) > 0
	}))
}
