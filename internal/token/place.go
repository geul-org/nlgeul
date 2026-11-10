//ff:func feature=token type=reader control=selection
//ff:what place — where an entry's form starts if it ends at j, and the structural facts of that place for its graph
package token

import "github.com/geul-org/koparse/pkg/rule"

// A piece with a lead sits in the stem's last syllable, so the stem
// ends where the form starts. more tells that 요 follows. A fused 어/아 that does
// not show (가, 서, 내, 세) is read when more of the tail follows, and
// also at the unit's end, where the stem under it is what a key table
// is asked for; inside a unit nothing would show the boundary, except
// right before an auxiliary verb's stem (가보다 = 가+아 보다). A fused ㅆ is read only on the vowels 었/았 fuses into.
func (sc *scan) place(e *entry, j int, more bool) (int, rule.Piece, bool) {
	f := []rune(e.form)
	start := j - len(f)
	if start < 0 || string(sc.s[start:j]) != e.form {
		return 0, rule.Piece{}, false
	}
	p := rule.Piece{Form: e.form, HasStem: start >= 1, Slot: e.slot, Next: -1}
	switch {
	case start == 0:
		return start, p, e.lead == 0
	case e.lead == 0:
		p.Prev = sc.s[start-1]
		return start, p, true
	case e.lead == leadVowel:
		return start, p, rule.FusedVowel(sc.s[start-1], len(f) > 0 || more || j == len(sc.s) || sc.auxNext(j))
	case e.lead == jongSS:
		return start, p, rule.FusedPast(sc.s[start-1])
	}
	base, ok := under(sc.s[start-1], e.lead)
	p.Base = string(base)
	return start, p, ok && (len(f) == 0 || rule.JongOf(sc.s[start-1]) == e.lead)
}
