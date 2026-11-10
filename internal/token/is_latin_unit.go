//ff:func feature=token type=helper control=sequence
//ff:what isLatinUnit — is a text one of the units of the table that are written in Latin letters (a run of digits and one of them is a number and its unit)
package token

import (
	"slices"
	"unicode"
	"unicode/utf8"
)

func isLatinUnit(text string) bool {
	return text != "" && text[0] < utf8.RuneSelf && unicode.IsLetter(rune(text[0])) && slices.Contains(unitWords, text)
}
