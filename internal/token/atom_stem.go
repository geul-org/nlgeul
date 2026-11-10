//ff:func feature=token type=reader control=sequence
//ff:what atomStem — when the stem from b to end is exactly one atom: its name (the readings that join it, or the run's own) and its verdict
package token

func (sc *scan) atomStem(b, end int) (string, float64) {
	a, ok := sc.atoms[b]
	if !ok || a.end != end {
		return "", 0
	}
	if a.names == "" {
		return runName, sc.ix.run
	}
	if a.num {
		return a.names, sc.ix.number
	}
	return a.names, 1
}
