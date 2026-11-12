//ff:func feature=token type=reader control=sequence
//ff:what stemOf — the stem of a word as it is printed and looked up, its name and its verdict (0 for a new word): the letters before the tail, with the last syllable as it was before a piece was fused into it, the word a fused form stands for, or known words bound by a slash or a hyphen; sole tells that the word is the whole of its unit, which the stem's graph is told
package token

// A fused form's word is given by the closed set, so it is known under
// the form's name when no table has it.
func (sc *scan) stemOf(b int, t tail, sole bool) (string, string, float64) {
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	if t.as != "" {
		return stem, t.as, sc.ix.number
	}
	name, v := sc.ix.known(stem, tailClass(t.pieces), sole)
	if v <= 0 {
		name, v = sc.atomStem(b, t.start)
	}
	if v <= 0 && t.label != "" {
		name, v = t.label, 1
	}
	if v <= 0 {
		name, v = sc.bound(stem, tailClass(t.pieces))
	}
	return stem, name, v
}
