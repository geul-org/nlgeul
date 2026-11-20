//ff:func feature=parse type=helper control=sequence
//ff:what isLatinUnit — is a text one of the units of the table that are written in Latin letters (a run of digits and one of them is a number and its unit)
package parse

import (
	"slices"
)

func isLatinUnit(text string) bool {
	return text != "" && isLatinLetter(firstRune(text)) && slices.Contains(unitWords, text)
}
