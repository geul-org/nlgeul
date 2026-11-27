//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what romanLetters — is a text made of the Roman numeral letters of Unicode alone (romanRange: Ⅱ, ⅳ, ⅩⅢ), each of which is a Roman numeral by itself; false for an empty text
package parse

func romanLetters(text string) bool {
	for _, r := range text {
		if r < romanRange[0] || r > romanRange[1] {
			return false
		}
	}
	return text != ""
}
