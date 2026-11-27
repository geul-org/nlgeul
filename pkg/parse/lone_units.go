//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loneUnits — the candidates of a stem whose letters are a unit of the unit table (m, km, kg, 년, 도), read with no number written onto it: one for each line of the table with those letters, named as the table names it (단위:미터), with the verdict of the graph loneUnitGraph, a new word's, and the kind that stands behind the words of the word table (kindBehind: after a number written apart the letters may be its unit, 700 m, 사십 년); a unit written in Hangul syllables is read so only with no piece on it (년, not 년 before 을: its letters are a nominal of the word table, and with a piece on it the candidate turned a first candidate and pushed right ones out of the first three on the answer sheet)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (ix *index) loneUnits(stem string, t tail) []kindCandidate {
	var out []kindCandidate
	if len(t.pieces) > 0 && slices.ContainsFunc([]rune(stem), func(r rune) bool { return rule.JongOf(r) >= 0 }) {
		return nil
	}
	for _, row := range unitRows {
		if row.Cells[0] != stem {
			continue
		}
		out = append(out, kindCandidate{kindBehind, unitLabel + row.Cells[1], ix.loneUnit})
	}
	return out
}
