//ff:func feature=token type=reader control=iteration dimension=1
//ff:what crossesChange — does the new word from b before tail t (which ends at j) stand over a place inside its stem where a run of Latin letters and digits meets a Hangul syllable (change), such that the stem from b to that place is one atom or a known word with no marker, and the letters from that place to j read as a known word or a run (readsKnown: A씨, K팝, 3분); a tail that names its own stem or stands for a number is not asked
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The stem ends where the tail starts, or one syllable before when a
// piece is fused into its last syllable. The place may be either way
// round (a run before Hangul or Hangul before a run), and a run of
// digits counts as a run.
func (sc *scan) crossesChange(b, j int, t tail) bool {
	if t.only != "" || t.as != "" {
		return false
	}
	end := t.start
	if t.base != 0 {
		end = t.start - 1
	}
	for k := b + 1; k < end; k++ {
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
