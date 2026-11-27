//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what approxAt — the approximate numbers that end at a place, each as its two numbers (approxPair: 육 and 칠, 이 and 삼십): every second number that ends there, however long (as far back as the letters of numerals go), with every first number that ends where it begins
package parse

func (sc *scan) approxAt(end int) [][2]string {
	var out [][2]string
	for n := 1; n < end && (numeralLetter(sc.s[end-n]) || string(sc.s[end-n]) == sinoSixAfter); n++ {
		second := string(sc.s[end-n : end])
		for _, first := range sc.numerals(end - n) {
			if approxPair(first.text, second) {
				out = append(out, [2]string{first.text, second})
			}
		}
	}
	return out
}
