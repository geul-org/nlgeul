//ff:func feature=mini type=helper control=sequence
//ff:what jungOf — the vowel index of a Hangul syllable (ㅏ 0 … ㅣ 20); -1 when r is not a Hangul syllable
package mini

func jungOf(r rune) int {
	if r < 0xAC00 || r > 0xD7A3 {
		return -1
	}
	return int(r-0xAC00) / 28 % 21
}
