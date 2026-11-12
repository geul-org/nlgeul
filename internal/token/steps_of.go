//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stepsOf — the candidates of one word: the stem from b (a known word, a known word written wrong, an atom or a new word) and the markers after it as the parts of a word, with its verdict and preference; one candidate for each way the stem stands (stemReads: each kind of a known word, each set of a closed-set word, each name of an atom or a sign); a stem no table lists stands as the new word's graph judges it, and a word whose stem does not stand has a verdict of 0 or less
package token

// A stem the word table has stands as its kind's graph judges it; the
// force of "rarely seen" is set so that one observation still stands
// above a new word. A stem that a piece is fused into
// is printed, and looked up, with its last syllable as it was before
// fusing (마신 → 마시). In a stem an atom counts as one token for
// preference. The boundary before an auxiliary verb's stem is given
// back only when the stem stands as a predicate (주 before the modifying
// 는, not 주 the noun before the particle 는).
func (sc *scan) stepsOf(b, end int, t tail) []step {
	var out []step
	for _, r := range sc.stemReads(b, t, b == 0 && end == len(sc.s)) {
		out = append(out, sc.stepOf(b, end, t, r))
	}
	return out
}
