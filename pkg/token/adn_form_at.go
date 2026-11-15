//ff:func feature=token type=reader control=iteration dimension=2
//ff:what adnFormAt — do the letters from b to k read as the modifying form of a predicate: some run of markers that ends at k and whose last piece is a modifying closing leaves, read after the letters from b (its last syllable restored where it may be), a stem that stands as a verb or an adjective of the tables
package token

import "strings"

func (sc *scan) adnFormAt(b, k int) bool {
	var ways []tail
	for _, t := range sc.tailsAt(k) {
		if t.start > b && len(t.pieces) > 0 && strings.HasPrefix(t.pieces[len(t.pieces)-1].names, adnPrefix) {
			ways = append(ways, sc.variants(b, t)...)
		}
	}
	for _, v := range ways {
		if !sc.restored(b, v) {
			continue
		}
		for _, r := range sc.stemCandidates(b, v, false) {
			if r.v > 0 && r.kind&(kindVerb|kindAdj) != 0 {
				return true
			}
		}
	}
	return false
}
