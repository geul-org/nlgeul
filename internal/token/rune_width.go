//ff:func feature=token type=helper control=sequence
//ff:what runeWidth — the cells a letter takes in a terminal: 0 for combining and zero width letters, 2 for East Asian wide letters (Hangul, kana, CJK, full width forms), 1 for the rest
package token

func runeWidth(r rune) int {
	if inRanges(r, zeroCells) {
		return 0
	}
	if inRanges(r, wideCells) {
		return 2
	}
	return 1
}
