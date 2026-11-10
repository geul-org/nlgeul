//ff:func feature=mini type=reader control=sequence
//ff:what prefinals — read the slots left of a closing in their order 시 < 었 < 겠 < 더 (endings.md section 1); where the stem ends, the length read, and whether a ㅆ fused into the stem was read
package mini

import "strings"

// Every slot read adds a syllable (2). 었/았/였 is read twice at most.
// A fused ㅆ belongs to the stem's last syllable: it adds 1 and reading
// stops there.
func prefinals(s []rune, start, size int) (int, int, bool) {
	i := start
	if i > 0 && s[i-1] == '더' {
		i, size = i-1, size+2
	}
	if i > 0 && s[i-1] == '겠' {
		i, size = i-1, size+2
	}
	if i > 0 && strings.ContainsRune("었았였", s[i-1]) {
		i, size = i-1, size+2
	}
	if i > 0 && strings.ContainsRune("었았였", s[i-1]) {
		i, size = i-1, size+2
	}
	if i > 0 && fusedPast(s[i-1]) {
		return i, size + 1, true
	}
	if i > 0 && s[i-1] == '시' {
		i, size = i-1, size+2
	}
	return i, size, false
}
