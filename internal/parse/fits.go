//ff:func feature=parse type=reader control=selection
//ff:what fits — does the syllable before a piece meet the piece's condition; a non-Hangul character meets every condition
package parse

func fits(prev rune, after int) bool {
	j := jongOf(prev)
	switch after {
	case afterBatchim:
		return j != 0
	case afterHard:
		return j != 0 && j != jongL
	case afterOpen:
		return j <= 0 || j == jongL
	case afterVowel:
		return j <= 0
	}
	return true
}
