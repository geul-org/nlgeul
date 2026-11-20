//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what crossesChange — does the new word from b before tail t (which ends at j) stand over a place inside its stem where a run of Latin letters and digits meets a Hangul syllable (change), such that the stem from b to that place is one atom or a known word with no marker, and the letters from that place to j read as a known word or a run (readsKnown: A씨, K팝, 3분), the place right before the stem's last syllable a piece is fused into among them (git했다); a tail that names its own stem or stands for a number is not asked
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// The stem ends where the tail starts. When a piece is fused into its
// last syllable the stem is the letters before that syllable and the
// syllable restored, so the place right before it is inside the stem
// too (git했다: git하). The place may be either way round (a run before
// Hangul or Hangul before a run), and a run of digits counts as a run.
func (sc *scan) crossesChange(b, j int, t tail) bool {
	if t.only != "" || t.as != "" {
		return false
	}
	for k := b + 1; k < t.start; k++ {
		if !sc.change(k) {
			continue
		}
		left := len(sc.atomStem(b, k)) > 0 || sc.ix.known(sc.letters(b, k), rule.TailBare, false) > 0
		if left && sc.readsKnown(k, j) {
			return true
		}
	}
	return false
}
