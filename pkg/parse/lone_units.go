//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what loneUnits — the candidates of a stem whose letters are a unit of the unit table written in letters that are no Hangul syllables (m, km, kg), read with no number written onto it: one for each line of the table with those letters, named as the table names it (단위:미터), with the verdict of the graph loneUnitGraph, a new word's (it stands behind the word of the word table: after a number written apart the letters may be its unit, 700 m)
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (ix *index) loneUnits(stem string) []kindCandidate {
	var out []kindCandidate
	if slices.ContainsFunc([]rune(stem), func(r rune) bool { return rule.JongOf(r) >= 0 }) {
		return nil
	}
	for _, row := range unitRows {
		if row.Cells[0] != stem {
			continue
		}
		out = append(out, kindCandidate{0, unitLabel + row.Cells[1], ix.loneUnit})
	}
	return out
}
