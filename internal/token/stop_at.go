//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stopAt — the full stop or the dots that end at j, read as a piece that closes the word before it: where it starts and the readings that stand
package token

// The piece costs nothing in preference: the tables give it no count,
// and it is not a syllable of a stem.
func (sc *scan) stopAt(j int) (piece, int, bool) {
	if j < 1 || sc.s[j-1] != '.' || (j < len(sc.s) && sc.s[j] == '.') {
		return piece{}, 0, false
	}
	start := j
	for start > 0 && sc.s[start-1] == '.' {
		start--
	}
	names, v := sc.judgeMark(sc.ix.stops, sc.stopFacts(start, j), string(sc.s[start:j]))
	return piece{text: string(sc.s[start:j]), names: names, verdict: v}, start, v > 0 && start > 0
}
