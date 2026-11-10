//ff:func feature=token type=reader control=selection
//ff:what place — where an entry's form starts if it ends at j, and the structural facts of that place for its graph
package token

import "github.com/geul-org/koparse/pkg/rule"

// more tells that 요 follows. A piece with a lead sits in the stem's
// last syllable, so the stem ends where the form starts.
func (sc *scan) place(e *entry, j int, more bool) (int, rule.Piece, bool) {
	f := []rune(e.form)
	start := j - len(f)
	if start < 0 || string(sc.s[start:j]) != e.form {
		return 0, rule.Piece{}, false
	}
	p := rule.Piece{Form: e.form, HasStem: start >= 1, PrevEnd: rule.EndOther, Slot: e.slot, Next: -1}
	switch {
	case start == 0:
		return start, p, e.lead == 0
	case e.lead == 0:
		p.PrevEnd = endClass(sc.s[start-1])
		return start, p, true
	case e.lead == leadVowel:
		return start, p, fusedVowel(sc.s[start-1], len(f) > 0 || more)
	}
	p.Base = string(sc.s[start-1] - rune(e.lead))
	return start, p, jongOf(sc.s[start-1]) == e.lead
}
