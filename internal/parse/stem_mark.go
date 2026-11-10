//ff:func feature=parse type=reader control=sequence
//ff:what stemMark — the marked text up to where the slots begin: the stem, with 셨 written as 시+었 when a stem stands before it, and the copula stem 이 as a piece of its own
package parse

// i is where the stem ends, fused tells that s[i-1] holds a fused ㅆ,
// cut that s[i-1] is the copula stem 이.
func stemMark(s []rune, i int, fused, cut bool) string {
	if cut {
		return affixMark(s[:i-1], false) + "+이"
	}
	if fused && s[i-1] == '셨' && i >= 2 {
		return affixMark(s[:i-1], false) + "+시+었"
	}
	return affixMark(s[:i], fused)
}
