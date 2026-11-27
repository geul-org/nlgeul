//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what withOrderLetter — the numbers before the unit 목 with the letter that numbers an item (orderLetter) added: a numeral of the same letter (사 of 사목) is dropped, as the letter stands before 목 as the letter alone
package parse

func withOrderLetter(words, letters []numeral) []numeral {
	var out []numeral
	for _, w := range words {
		if !hasNumeralText(letters, w.text) {
			out = append(out, w)
		}
	}
	return append(out, letters...)
}
