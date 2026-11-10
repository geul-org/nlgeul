//ff:func feature=parse type=reader control=sequence
//ff:what aMark — a main verb in its 어/아 form, marked: the 어/아 syllable as a piece (되+어), a fused one in one piece with the stem (작성+해)
package parse

import "strings"

func aMark(stem []rune) string {
	n := len(stem)
	if n >= 2 && strings.ContainsRune("어아여", stem[n-1]) {
		return affixMark(stem[:n-1], false) + "+" + string(stem[n-1])
	}
	return affixMark(stem, true)
}
