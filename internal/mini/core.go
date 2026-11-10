//ff:func feature=mini type=helper control=sequence
//ff:what core — the word without the punctuation and symbols around it
package mini

import (
	"strings"
	"unicode"
)

func core(word string) []rune {
	return []rune(strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	}))
}
