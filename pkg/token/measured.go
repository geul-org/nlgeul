//ff:func feature=token type=reader control=iteration dimension=1
//ff:what measured — the contests of a sign of a pair alone, when it stands right after a digit and is a sign of the unit table (3", 5'), with a branch added for each line of the table that has it: the sign as a unit (단위:인치), with the preference of a unit and the verdict of measureGraph (lowered where the sign may close a pair open before it: closesPair); after the candidates as a sign of a pair and the sign as its letter alone
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func (ix *index) measured(contests []Contest, u unit) []Contest {
	if !u.measure || len(contests) == 0 {
		return contests
	}
	m := rule.Mark{}
	if closesPair(u.candidates) {
		m.Facts = rule.MarkOpened
	}
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyMark, m)
	v, _ := evaluate(ix.measure, ctx, false)
	sign := string([]rune(u.text)[:u.head])
	for _, row := range unitRows {
		if row.Cells[0] != sign {
			continue
		}
		p := Part{Text: sign, Name: nameOf(unitLabel + row.Cells[1])}
		contests[0].Branches = append(contests[0].Branches, Branch{Words: []Eojeol{{Parts: []Part{p}, End: u.head}}, Verdict: v, Pref: ix.afterPrefs[false]})
	}
	return contests
}
