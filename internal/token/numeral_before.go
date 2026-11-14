//ff:func feature=token type=reader control=iteration dimension=1
//ff:what numeralBefore — do the letters right before place p read as a numeral of the number table, Sino-Korean or native
package token

func (sc *scan) numeralBefore(p int) bool {
	for a := p - 1; a >= 0; a-- {
		if t := string(sc.s[a:p]); sinoTokens(t) > 0 || nativeTokens(t) > 0 {
			return true
		}
	}
	return false
}
