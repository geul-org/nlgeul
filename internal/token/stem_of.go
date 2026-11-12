//ff:func feature=token type=reader control=sequence
//ff:what stemOf — the stem of a word as it is printed and looked up, its name and its verdict (0 for a new word): the letters before the tail, with the last syllable as it was before a piece was fused into it, the word a fused form stands for, or known words bound by a slash or a hyphen (boundTypo: a word of them written wrong makes the right form of the whole); a stem written wrong is the right form, judged as a misspelt form (misspelt) and by nothing else, and the letters it is written with come back as the fourth value (empty for a stem written right); sole tells that the word is the whole of its unit, which the stem's graph is told
package token

// A fused form's word is given by the closed set, so it is known under
// the form's name when no table has it.
func (sc *scan) stemOf(b int, t tail, sole bool) (string, string, float64, string) {
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	if t.as != "" {
		return stem, t.as, sc.ix.number, ""
	}
	if t.wrote != "" {
		name, v := sc.ix.misspelt(stem, t.kinds, tailClass(t.pieces), sole)
		return stem, name, v, t.wrote
	}
	name, v := sc.ix.known(stem, tailClass(t.pieces), sole)
	if v <= 0 {
		name, v = sc.atomStem(b, t.start)
	}
	if v <= 0 && t.label != "" {
		name, v = t.label, 1
	}
	if v > 0 {
		return stem, name, v, ""
	}
	right, name, v := sc.boundOf(stem, tailClass(t.pieces), t.boundTypo)
	if right != stem {
		return right, name, v, stem
	}
	return stem, name, v, ""
}
