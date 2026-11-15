//ff:func feature=token type=reader control=iteration dimension=2
//ff:what readsAs — do the letters from k to j read as a known word that the test given accepts, with no marker or with a run of markers that ends at j: some way a tail is read there (its last syllable restored where it may be) leaves a stem from k that stands (a verdict above 0) as the test wants it (its kind, its stem); a new word does not count
package token

func (sc *scan) readsAs(k, j int, ok func(stemCandidate) bool) bool {
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
		for _, r := range sc.stemCandidates(k, v, false) {
			if r.v > 0 && ok(r) {
				return true
			}
		}
	}
	return false
}
