//ff:func feature=token type=reader control=sequence
//ff:what stemCandidates — the candidates of the stem of a word as it is printed and looked up, each with its kind, name and verdict (one with no name and a verdict of 0 for a new word): the letters before the tail, with the last syllable as it was before a piece was fused into it, the word a fused form stands for, or known words bound by a slash or a hyphen (boundTypo: a word of them written wrong makes the right form of the whole); a stem no table lists is a new word only up to newStemMax tokens (farNew; the stem alone, and not when it is the whole of its unit), a stem written wrong is the right form, judged as a misspelt form (misspeltCandidates) and by nothing else, with the letters it is written with; a stem of the tables is read once for each kind and each set that lets it stand, an atom once for each name that joins it, a bound stem once for each candidate of its sign; sole tells that the word is the whole of its unit, which the stem's graph is told
package token

// A fused form's word is given by the closed set, so it is known under
// the form's name when no table has it.
func (sc *scan) stemCandidates(b int, t tail, sole bool) []stemCandidate {
	stem := sc.letters(b, t.start)
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	if t.as != "" {
		return []stemCandidate{{kindCandidate{0, t.as, sc.ix.number}, stem, ""}}
	}
	if t.wrote != "" {
		if candidates := stemWith(stem, t.wrote, sc.ix.misspeltCandidates(stem, t.kinds, tailClass(t.pieces), sole)); len(candidates) > 0 {
			return candidates
		}
		if sc.farNew(b, t, sole) {
			return nil
		}
		return []stemCandidate{{kindCandidate{}, stem, t.wrote}}
	}
	if candidates := stemWith(stem, "", sc.ix.knownCandidates(stem, tailClass(t.pieces), sole)); len(candidates) > 0 {
		return candidates
	}
	if candidates := stemWith(stem, "", kindCandidatesOf(sc.atomStem(b, t.start))); len(candidates) > 0 {
		return candidates
	}
	if t.label != "" {
		return []stemCandidate{{kindCandidate{0, t.label, 1}, stem, ""}}
	}
	right, marks := sc.boundOf(stem, tailClass(t.pieces), t.boundTypo)
	wrote := ""
	if right != stem {
		wrote = stem
	}
	if candidates := stemWith(right, wrote, kindCandidatesOf(marks)); len(candidates) > 0 {
		return candidates
	}
	if sc.farNew(b, t, sole) {
		return nil
	}
	return []stemCandidate{{kindCandidate{}, stem, ""}}
}
