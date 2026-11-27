//ff:func feature=parse type=helper control=selection
//ff:what numberScript — the form of a number written in the letter given: 한글십진 for a Hangul syllable, 한자십진 for a Chinese character (or 〇), 로마십진 for a Roman numeral letter of Unicode (the shape 로마 숫자 of the tables of special letters), 십진 for a digit; none (empty) for any other letter
package parse

func numberScript(r rune) string {
	switch {
	case r >= '0' && r <= '9':
		return decimalForm
	case r >= 0xAC00 && r <= 0xD7A3:
		return hangulForm
	case r >= 0x4E00 && r <= 0x9FFF, r >= 0x3400 && r <= 0x4DBF, r == '〇':
		return hanjaForm
	case formShaped(r, romanShape):
		return romanForm
	}
	return ""
}
