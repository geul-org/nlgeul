//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what stepsOf — the candidates of one word: the stem from b (a known word, a known word written wrong, an atom or a new word) and the markers after it as the parts of a word, with its verdict and preference; one candidate for each way the stem stands (stemCandidates: each kind of a known word, each set of a closed-set word, each name of an atom or a sign), and one more for a family name before a title (surnameCandidate); a stem no table lists stands as the new word's graph judges it, and a word whose stem does not stand has a verdict of 0 or less; a candidate that stands behind the words of the word table (kindBehind) is left out where the new word of its letters does not stand (stepOf gives it no verdict there)
package parse

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
	for _, r := range append(sc.stemCandidates(b, t, b == 0 && end == len(sc.s), sc.mainBeforeAux(b, end, t)), sc.surnameCandidate(b, t)...) {
		e := sc.stepOf(b, end, t, r)
		if r.kind == kindBehind && e.verdict <= 0 {
			continue
		}
		out = append(out, e)
	}
	return out
}
