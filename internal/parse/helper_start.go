//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what helperStart — where an auxiliary verb's stem begins at the end of stem, when the 어/아 form of a main verb stands before it (해주, 해봐, 작성해준); -1 when there is none
package parse

// fused tells that a tail piece is fused into the stem's last syllable.
func helperStart(stem []rune, fused bool) int {
	plain := append([]rune(nil), stem...)
	if fused {
		plain[len(plain)-1] = unfuse(plain[len(plain)-1])
	}
	for _, h := range helperStems {
		k, ok := suffixAt(plain, len(plain), h)
		if ok && aForm(stem[:k]) {
			return k
		}
	}
	return -1
}
