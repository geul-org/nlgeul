//ff:func feature=token type=reader control=sequence
//ff:what honors — the readings of a letter as the honorific 시 with a piece fused under it: the stem ends before the letter, or before the link vowel 으 that stands before it (linked)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// start is the place after the letter, p the piece fused under it;
// bare tells that the piece is the final consonant alone, so that what
// follows the letter is looked at.
func (sc *scan) honors(p piece, start int, bare bool) []tail {
	at := start - 1
	e, gate := sc.ix.honorOf(sc.s[at])
	if e == nil {
		return nil
	}
	facts := rule.Piece{Form: "시", HasStem: at >= 1, Next: -1, Right: rule.RightDone}
	if at >= 1 {
		facts.Prev = sc.s[at-1]
	}
	if bare {
		facts.Right = sc.right(start)
	}
	var out []tail
	if v := sc.judge(e, at, start, facts); v > 0 {
		h := piece{"시", honorName, v, sc.ix.pref(honorCount, "시"), rule.TailPred, false}
		out = append(out, tail{start: at, pieces: []piece{h, p}, gate: gate, head: headEuOpen | headS})
		out = append(out, withGate(sc.linked(at, []piece{h, p}, honorName, false), gate)...)
	}
	return out
}
