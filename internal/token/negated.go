//ff:func feature=token type=reader control=iteration dimension=1
//ff:what negated — the readings of 잖 before a closing: 지+않, the closing attached to 않, alone or after the slots before 지 (했잖아 → 하+(였)+지+(않)+아)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// The tails are those of the closings read after the letter. A tail
// fused into the letter, or one of the copula, is not such a closing.
func (sc *scan) negated(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		at := t.start - 1
		if t.base != 0 || t.only != "" || t.irr != 0 || at < 0 || sc.s[at] != negLetter {
			continue
		}
		facts := rule.Piece{Form: string(negLetter), HasStem: at >= 1, Next: rule.SlotClosing}
		v := sc.judge(sc.ix.neg, at, t.start, facts)
		if v <= 0 {
			continue
		}
		ji := piece{"지", jiName, v, sc.ix.pref(sc.ix.ji.count, "지"), rule.TailPred, false}
		ps := append(partsOf(ji, negParts), t.pieces...)
		out = append(out, tail{start: at, pieces: ps, noYo: t.noYo, close: sc.ix.ji})
		for _, l := range joined(sc.slots(at, rule.SlotClosing), ps...) {
			l.noYo, l.close = t.noYo, sc.ix.ji
			out = append(out, l)
		}
	}
	return out
}
