//ff:func feature=token type=reader control=sequence
//ff:what stemReads — the candidates of the stem of a word as it is printed and looked up, each with its kind, name and verdict (one with no name and a verdict of 0 for a new word): the letters before the tail, with the last syllable as it was before a piece was fused into it, the word a fused form stands for, or known words bound by a slash or a hyphen (boundTypo: a word of them written wrong makes the right form of the whole); a stem written wrong is the right form, judged as a misspelt form (misspeltReads) and by nothing else, with the letters it is written with; a stem of the tables is read once for each kind and each set that lets it stand, an atom once for each name that joins it, a bound stem once for each candidate of its sign; sole tells that the word is the whole of its unit, which the stem's graph is told
package token

// A fused form's word is given by the closed set, so it is known under
// the form's name when no table has it.
func (sc *scan) stemReads(b int, t tail, sole bool) []stemRead {
	stem := string(sc.s[b:t.start])
	if t.base != 0 {
		stem = string(sc.s[b:t.start-1]) + string(t.base)
	}
	if t.shown != "" {
		stem = t.shown
	}
	if t.as != "" {
		return []stemRead{{kindRead{0, t.as, sc.ix.number}, stem, ""}}
	}
	if t.wrote != "" {
		if reads := stemWith(stem, t.wrote, sc.ix.misspeltReads(stem, t.kinds, tailClass(t.pieces), sole)); len(reads) > 0 {
			return reads
		}
		return []stemRead{{kindRead{}, stem, t.wrote}}
	}
	if reads := stemWith(stem, "", sc.ix.knownReads(stem, tailClass(t.pieces), sole)); len(reads) > 0 {
		return reads
	}
	if reads := stemWith(stem, "", kindReadsOf(sc.atomStem(b, t.start))); len(reads) > 0 {
		return reads
	}
	if t.label != "" {
		return []stemRead{{kindRead{0, t.label, 1}, stem, ""}}
	}
	right, marks := sc.boundOf(stem, tailClass(t.pieces), t.boundTypo)
	wrote := ""
	if right != stem {
		wrote = stem
	}
	if reads := stemWith(right, wrote, kindReadsOf(marks)); len(reads) > 0 {
		return reads
	}
	return []stemRead{{kindRead{}, stem, ""}}
}
