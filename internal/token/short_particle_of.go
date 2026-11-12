//ff:func feature=token type=reader control=sequence
//ff:what shortParticleOf — one particle shortened with 는 read from a line of its table: its form, what must stand before it, its case and the form it is short for; it is named with its case, its form and the form it is short for
package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func shortParticleOf(row data.Row) (entry, error) {
	lead, form, err := parse.FormOf(row, 0)
	if err != nil {
		return entry{}, err
	}
	if lead != 0 || form == "" {
		return entry{}, row.Errf(0, "조사의 꼴은 한글 음절로만 적는다")
	}
	after, err := row.Pick(1, parse.AfterNames)
	if err != nil {
		return entry{}, err
	}
	kase, err := row.Pick(2, parse.CaseNames)
	if err != nil {
		return entry{}, err
	}
	name := row.Cells[2] + ":" + form + "(" + row.Cells[3] + ")"
	return entry{name: name, form: form, class: classParticle, after: after, slot: rule.LinkTopic, kase: kase}, nil
}
