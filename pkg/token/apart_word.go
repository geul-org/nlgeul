//ff:func feature=token type=reader control=iteration dimension=2
//ff:what apartWord — do the letters from k to j, read as a known word with no marker or with a run of markers that ends at j (each way a tail is read, its last syllable restored where it may be), read only as a word the spacing articles write apart: every reading's stem ends at one place, the stem's letters are a title of article 48 or a dependent noun of article 42 (titleWords, afterNounWords), and no reading is a verb's or an adjective's
package token

import "slices"

// The rule leaves nothing else to read there: when the same letters
// read as another nominal or as a predicate (양의, 내는), or are cut in
// two ways, the articles do not tell between them.
func (sc *scan) apartWord(k, j int) bool {
	var ways []tail
	for _, t := range append([]tail{{start: j}}, sc.tailsAt(j)...) {
		if t.start > k {
			ways = append(ways, sc.variants(k, t)...)
		}
	}
	ends, kinds, apart := map[int]bool{}, 0, false
	for _, v := range ways {
		if !sc.restored(k, v) {
			continue
		}
		for _, r := range sc.stemCandidates(k, v, false) {
			if r.v <= 0 {
				continue
			}
			stem := sc.letters(k, v.start)
			ends[v.start], kinds = true, kinds|r.kind
			apart = apart || (v.base == 0 && v.shown == "" && (slices.Contains(titleWords, stem) || slices.Contains(afterNounWords, stem)))
		}
	}
	return apart && len(ends) == 1 && kinds&(kindVerb|kindAdj) == 0
}
