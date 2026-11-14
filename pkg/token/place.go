//ff:func feature=token type=reader control=selection
//ff:what place — where an entry's form starts if it ends at j, and the structural facts of that place for its graph (a closing written wrong tells its graph that it is a misspelt form); a fused 어/아 or ㅆ is read on the syllables it fuses into, and after a syllable of unspelt where the writing leaves the 어 out
package token

import (
	"unicode/utf8"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A piece with a lead sits in the stem's last syllable, so the stem
// ends where the form starts. more tells that 요 follows. A fused 어/아 that does
// not show (가, 서, 내, 세) is read when more of the tail follows, and
// also at the unit's end, where the stem under it is what a key table
// is asked for; inside a unit nothing would show the boundary, except
// right before an auxiliary verb's stem (가보다 = 가+아 보다). A fused ㅆ is read only on the vowels 었/았 fuses into. After a
// syllable of unspelt (되) both are read as well, with an 어 the
// writing leaves out (되도, 됬다).
func (sc *scan) place(e *entry, j int, more bool) (int, rule.Piece, bool) {
	n := utf8.RuneCountInString(e.form)
	start := j - n
	if start < 0 || matchAt(sc.s, start, e.form) != j {
		return 0, rule.Piece{}, false
	}
	p := rule.Piece{Form: e.form, HasStem: start >= 1 || sc.attach, Slot: e.slot, Next: -1, Misspelt: e.wrote != ""}
	switch {
	case start == 0:
		return start, p, e.lead == 0
	case e.lead == 0:
		p.Prev = sc.s[start-1]
		return start, p, true
	case e.lead == rule.LeadVowel:
		open := n > 0 || more || j == len(sc.s) || sc.auxNext(j)
		return start, p, rule.FusedVowel(sc.s[start-1], open) || unspeltAt(sc.s[start-1], e.lead, open)
	case e.lead == rule.JongSS:
		return start, p, rule.FusedPast(sc.s[start-1]) || unspeltAt(sc.s[start-1], e.lead, false)
	}
	base, ok := under(sc.s[start-1], e.lead)
	p.Base = string(base)
	return start, p, ok && (n == 0 || rule.JongOf(sc.s[start-1]) == e.lead)
}
