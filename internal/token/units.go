//ff:func feature=token type=reader control=iteration dimension=1
//ff:what units — the pieces after a number that end at j: 여, a unit of the table, 째, in that order, the first of them right after a number
package token

import "github.com/geul-org/koparse/pkg/rule"

// next is the place of the piece to the right (0 for none). A piece
// next to a run that is no number is judged too, so that the trace
// tells why it is no unit.
func (sc *scan) units(j, next int) []tail {
	var out []tail
	for _, e := range sc.ix.after[sc.at(j-1)] {
		start := j - len([]rune(e.form))
		if start < 1 || string(sc.s[start:j]) != e.form || (next != 0 && e.slot >= next) {
			continue
		}
		lefts := sc.units(start, e.slot)
		if len(lefts) == 0 && !rule.RunRune(sc.s[start-1]) {
			continue
		}
		v := sc.judge(e, start, j, rule.Piece{Form: e.form, HasStem: true, Next: -1, Num: sc.numEnds[start] || len(lefts) > 0})
		if v <= 0 {
			continue
		}
		pc := piece{e.form, e.name, v, sc.ix.afterPref, rule.TailAux, false}
		if sc.numEnds[start] {
			out = append(out, tail{start: start, pieces: []piece{pc}})
		}
		out = append(out, joined(lefts, pc)...)
	}
	return out
}
