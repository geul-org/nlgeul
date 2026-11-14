//ff:func feature=token type=reader control=iteration dimension=1
//ff:what signWords — the candidates that end at j with a sign between letters (sc.mid) read as a word of its own: each candidate kept before the sign (keptAt), grown by the sign alone, with its preference 0 as a sign at a unit's edge has; the sign is named 기호 until the stretches are made (expandMid)
package token

func (sc *scan) signWords(cands []*path, best [][]*path, j int) []*path {
	q := j - 1
	if _, ok := sc.mid[q]; !ok || (sc.starts != nil && !sc.starts[q]) {
		return cands
	}
	e := sc.wordOf(q, j, tail{start: j, as: midName})
	e.pref = 0
	for _, p := range sc.keptAt(best, q) {
		cands = append(cands, grow(p, e))
	}
	return cands
}
