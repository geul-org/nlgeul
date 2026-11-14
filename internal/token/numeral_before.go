//ff:func feature=token type=reader control=iteration dimension=1
//ff:what numeralBefore — do the letters right before place p read as a numeral of the number table, Sino-Korean or native
package token

// Only the letters a numeral can be made of are looked at, and no more
// of them than the longest one: sinoTokens of a text is not 0 only when
// each letter of it is a digit or a greater number (수 only at the front
// of a text of more than one letter), so the last letter of it alone,
// which is no 수, is not 0 either, and the longest text that can give
// a native numeral is numeralReach letters. Looking at the letters one
// after another back to the unit's head gives the same answer
// (numeralBeforeAll in the tests), but costs the square of the unit's
// length at each place.
func (sc *scan) numeralBefore(p int) bool {
	for a := p - 1; a >= 0 && a >= p-numeralReachMax; a-- {
		if t := string(sc.s[a:p]); sinoTokens(t) > 0 || nativeTokens(t) > 0 {
			return true
		}
	}
	return false
}
