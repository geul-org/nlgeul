//ff:func feature=token type=reader control=iteration dimension=1
//ff:what units — the pieces after a number that end at j: a greater number, 여, a unit of the table, 째, in that order, the first of them right after a number (a run of digits; for all but a greater number also Hangul numerals or a Latin letter of the table)
package token

import "github.com/geul-org/koparse/pkg/rule"

// next is the place of the piece to the right (0 for none). open tells
// that nothing is known to follow at j, so that letters running on are
// looked for. A piece next to a run that is no number is judged too,
// so that the trace tells why it is no unit.
func (sc *scan) units(j, next int, open bool) []tail {
	var out []tail
	for _, e := range sc.ix.after[sc.at(j-1)] {
		start := j - len([]rune(e.form))
		if start < 1 || string(sc.s[start:j]) != e.form || !placed(e.slot, next) {
			continue
		}
		lefts, words := sc.units(start, e.slot, false), sc.numerals(start)
		if e.slot == placeBig {
			words = nil
		}
		if len(lefts) == 0 && len(words) == 0 && !rule.RunRune(sc.s[start-1]) {
			continue
		}
		facts := rule.Piece{Form: e.form, HasStem: true, Next: -1, Right: rule.RightDone, Num: sc.numEnds[start] || len(lefts)+len(words) > 0}
		if open && sideOf(string(sc.at(j))) != 0 {
			facts.Right = rule.RightOther
		}
		v := sc.judge(e, start, j, facts)
		if v <= 0 {
			continue
		}
		pc := piece{e.form, e.name, v, sc.ix.afterPrefs[e.slot == placeBig], rule.TailAux, false}
		if sc.numEnds[start] {
			out = append(out, tail{start: start, pieces: []piece{pc}})
		}
		for _, w := range words {
			out = append(out, tail{start: start, pieces: []piece{pc}, only: w.text, as: numName, weight: w.pref})
		}
		out = append(out, joined(lefts, pc)...)
	}
	return out
}
