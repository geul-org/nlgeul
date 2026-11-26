//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what bareUnits — the tails of a suffix of bareSuffixes (당) that ends at j right after a unit of the unit table with no number before the unit (㎡당): the unit is the stem, named as the table names it; only a unit written in letters that are no Hangul syllables is read so (a unit's Hangul name is a nominal of the word table), and the suffix's graph is told that what it counts by stands before it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

// A unit right after a number is read by units; here no number ends
// where the unit begins. Each unit is taken as one of the table's
// entries, as after a number.
func (sc *scan) bareUnits(e *entry, start, j int) []tail {
	if !slices.Contains(bareSuffixes, e.form) || e.slot != placeOrder {
		return nil
	}
	var out []tail
	for _, u := range sc.ix.after[sc.s[start-1]] {
		b := start - len([]rune(u.form))
		if u.slot != placeUnit || b < 0 || string(sc.s[b:start]) != u.form || rule.JongOf(sc.s[start-1]) >= 0 || sc.numEnds[b] || len(sc.numerals(b)) > 0 {
			continue
		}
		v := sc.judge(e, start, j, rule.Piece{Form: e.form, HasStem: true, Next: -1, Right: rule.RightDone, Num: true})
		if v <= 0 {
			continue
		}
		pc := piece{e.form, e.name, v, sc.ix.afterPrefs[false], rule.TailAux, false, ""}
		out = append(out, tail{start: start, pieces: []piece{pc}, only: u.form, as: u.name, weight: sc.ix.afterPrefs[false]})
	}
	return out
}
