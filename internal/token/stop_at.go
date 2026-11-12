//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stopAt — the full stop, the dots or the comma that end at j inside a unit, read as a piece that closes the word before it (가,나 → 가+,[쉼표] 나): where it starts and a piece for each candidate that stands, with that candidate's verdict
package token

import "strings"

// The piece costs nothing in preference: the tables give it no count,
// and it is not a syllable of a stem. Something must stand before it in
// the unit: a sign at a unit's edge is sent up on its own (signBundles).
func (sc *scan) stopAt(j int) ([]piece, int) {
	if j < 1 || !strings.ContainsRune(stopSigns, sc.s[j-1]) || (j < len(sc.s) && sc.s[j] == sc.s[j-1]) {
		return nil, 0
	}
	start := j
	for start > 0 && sc.s[start-1] == sc.s[j-1] {
		start--
	}
	if start == 0 {
		return nil, 0
	}
	var out []piece
	for _, r := range sc.judgeMark(sc.ix.marks[sc.s[j-1]], sc.stopFacts(start, j), string(sc.s[start:j])) {
		out = append(out, piece{text: string(sc.s[start:j]), names: r.name, verdict: r.v})
	}
	return out, start
}
