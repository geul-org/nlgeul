//ff:func feature=parse type=reader control=sequence
//ff:what readsCopula — is the copula stem 이 read at the stem's end (morphemes.md section 2), and the length it adds
package parse

// The stem 이 is read when the closing holds it (예요), when the
// syllable before the tail is 이 with a word before it (2 more), or
// when the leftmost slot is 였 not after 하 (이 + 었; after 하 it is
// 하다's own 였). Closings that do not attach to 이 never read it.
func readsCopula(s []rune, i int, e ending) (bool, int) {
	if e.flags&fNoIda != 0 {
		return false, 0
	}
	if e.flags&fSelfIda != 0 {
		return true, 0
	}
	if i >= 2 && s[i-1] == '이' {
		return true, 2
	}
	return s[i] == '였' && s[i-1] != '하', 0
}
