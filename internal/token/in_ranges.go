//ff:func feature=token type=helper control=iteration dimension=1
//ff:what inRanges — whether the letter is in one of the ranges
package token

func inRanges(r rune, rs []runeRange) bool {
	for _, x := range rs {
		if r >= x.Lo && r <= x.Hi {
			return true
		}
	}
	return false
}
