//ff:func feature=parse type=helper control=sequence
//ff:what jongOf — the final-consonant index of a Hangul syllable (0 for none); -1 when r is not a Hangul syllable
package parse

func jongOf(r rune) int {
	if r < 0xAC00 || r > 0xD7A3 {
		return -1
	}
	return int(r-0xAC00) % 28
}
