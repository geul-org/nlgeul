//ff:func feature=token type=reader control=iteration dimension=1
//ff:what signWords — the candidates that end at j with a sign between letters (sc.mid) read as a word of its own, the whole run of that sign when it stands two or more in a row (midStart: ——, and only where the run ends): each candidate kept before the sign (keptAt), grown by the sign alone, with its preference 0 as a sign at a unit's edge has; the sign is named 기호 until the stretches are made (expandMid)
package token

func (sc *scan) signWords(cands []*path, best [][]*path, j int) []*path {
	q := sc.midStart(j)
	if _, ok := sc.mid[q]; !ok || (j < len(sc.s) && sc.s[j] == sc.s[j-1]) || (sc.starts != nil && !sc.starts[q]) {
		return cands
	}
	e := sc.wordOf(q, j, tail{start: j, as: midName})
	e.pref = 0
	for _, p := range sc.keptAt(best, q) {
		cands = append(cands, grow(p, e))
	}
	return cands
}
