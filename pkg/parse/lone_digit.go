//ff:func feature=parse type=reader control=sequence
//ff:what loneDigit — is a number before a unit a lone Sino-Korean digit (일 … 구) read with no word table: nothing then tells it from the first syllable of a common word
package parse

import "strings"

func (ix *index) loneDigit(numeral string) bool {
	return ix.layers == layerNone && len([]rune(numeral)) == 1 && strings.Contains(sinoDigits, numeral)
}
