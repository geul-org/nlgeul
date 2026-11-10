//ff:func feature=token type=reader control=sequence
//ff:what honors — the readings of a letter as the honorific 시 with a piece fused under it: the stem ends before the letter, or before the link vowel 으 that stands before it
package token

import "github.com/geul-org/koparse/pkg/rule"

// start is the place after the letter, p the piece fused under it;
// bare tells that the piece is the final consonant alone, so that what
// follows the letter is looked at.
func (sc *scan) honors(p piece, start int, bare bool) []tail {
	at := start - 1
	e, gate := sc.ix.honorOf(sc.s[at])
	if e == nil {
		return nil
	}
	h := piece{"시", honorName, 1, sc.ix.pref(honorCount, "시"), rule.TailPred, false}
	facts := rule.Piece{Form: "시", HasStem: at >= 1, Next: -1, Right: rule.RightDone}
	if at >= 1 {
		facts.Prev = sc.s[at-1]
	}
	if bare {
		facts.Right = sc.right(start)
	}
	var out []tail
	if sc.judge(e, at, start, facts) > 0 {
		out = append(out, tail{start: at, pieces: []piece{h, p}, gate: gate})
	}
	if at >= 2 && sc.s[at-1] == '으' && sc.judge(sc.ix.eu, at-1, at, rule.Piece{Form: "으", HasStem: true, Prev: sc.s[at-2], Next: -1}) > 0 {
		eu := piece{"으", euName, 1, sc.ix.pref(0, "으"), rule.TailPred, false}
		out = append(out, tail{start: at - 1, pieces: []piece{eu, h, p}, gate: gate})
	}
	return out
}
