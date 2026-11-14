//ff:func feature=token type=helper control=iteration dimension=1
//ff:what numeralBeforeAll — numeralBefore as it was before it looked back only as far as a numeral reaches: the letters right before place p up to the unit's head, each text of them asked whether it is a numeral of the number table (the answer that the short one has to give again)
package token

func numeralBeforeAll(sc *scan, p int) bool {
	for a := p - 1; a >= 0; a-- {
		if t := string(sc.s[a:p]); sinoTokens(t) > 0 || nativeTokens(t) > 0 {
			return true
		}
	}
	return false
}
