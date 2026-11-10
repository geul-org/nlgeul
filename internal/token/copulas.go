//ff:func feature=token type=reader control=iteration dimension=1
//ff:what copulas — the copula readings of some predicate tails: each tail that may stand after the stem 이 (said or dropped), judged by the copula's graph, with the stem as a piece before the tail
package token

import "github.com/geul-org/koparse/pkg/rule"

// A dropped stem is written (이) and leaves no trace in the letters.
func (sc *scan) copulas(tails []tail) []tail {
	var out []tail
	for _, t := range tails {
		start, e, ok := sc.copulaOf(t)
		if !ok {
			continue
		}
		p := rule.Piece{Form: "이", HasStem: start >= 1, Next: rule.SlotClosing, After: t.close.name}
		base := rune(0)
		if start >= 1 {
			p.Prev = sc.s[start-1]
		}
		if e == sc.ix.dropped && t.base != 0 {
			base, p.Prev = t.base, t.base
		}
		if sc.judge(e, start, start+len([]rune(e.form)), p) <= 0 {
			continue
		}
		pc := piece{"이", e.name, 1, sc.ix.pref(e.count, "이"), rule.TailCopula, e == sc.ix.dropped}
		if pc.unseen {
			pc.text = "(이)"
		}
		read := tail{start: start, pieces: append([]piece{pc}, t.pieces...), base: base, noYo: t.noYo, close: t.close}
		out = append(out, read)
		out = append(out, sc.afterUnits(read)...)
	}
	return out
}
