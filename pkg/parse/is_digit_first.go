//ff:func feature=parse type=helper control=sequence
//ff:what isDigitFirst — does a text begin with a digit 0 to 9
package parse

func isDigitFirst(text string) bool {
	return text != "" && text[0] >= '0' && text[0] <= '9'
}
