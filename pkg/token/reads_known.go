//ff:func feature=token type=reader control=iteration dimension=2
//ff:what readsKnown — do the letters from k to j read as a known word or a run of Latin letters and digits, with no marker or with a run of markers that ends at j: some way a tail is read there (its last syllable restored where it may be) leaves a stem from k that stands as a word of the tables, an atom, or another candidate with a verdict above 0 (a new word does not count)
package token

func (sc *scan) readsKnown(k, j int) bool {
	var ways []tail
	for _, t := range append([]tail{{start: j}}, sc.tailsAt(j)...) {
		if t.start > k {
			ways = append(ways, sc.variants(k, t)...)
		}
	}
	for _, v := range ways {
		if !sc.restored(k, v) {
			continue
		}
		for _, r := range sc.stemCandidates(k, v, false, false) {
			if r.v > 0 {
				return true
			}
		}
	}
	return false
}
