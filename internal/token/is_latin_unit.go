//ff:func feature=token type=helper control=sequence
//ff:what isLatinUnit — is a text one of the units written in Latin letters
package token

import "slices"

func isLatinUnit(text string) bool {
	return slices.Contains(latinUnits, text)
}
