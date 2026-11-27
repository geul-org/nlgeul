//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what romanLetters — is a text made of the Roman numeral letters of Unicode alone (the letters the tables of special letters give the shape 로마 숫자: Ⅱ, ⅳ, ⅩⅢ), each of which is a Roman numeral by itself; false for an empty text
package parse

func romanLetters(text string) bool {
	for _, r := range text {
		if !formShaped(r, romanShape) {
			return false
		}
	}
	return text != ""
}
