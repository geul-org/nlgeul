//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what hasNumeralText — does one of the numbers have the text given
package parse

func hasNumeralText(words []numeral, text string) bool {
	for _, w := range words {
		if w.text == text {
			return true
		}
	}
	return false
}
