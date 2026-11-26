//ff:func feature=parse type=reader control=sequence
//ff:what spokenOf — one marker written as it is spoken read from a line of its table: the letters it is written with (whole syllables, under a final consonant, or after a fused 어/아), the name of the standard marker, and the condition (none, or only after the copula stem)
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func spokenOf(row data.Row) (spoken, error) {
	lead, form, err := FormOf(row, 0)
	if err != nil {
		return spoken{}, err
	}
	if form == "" {
		return spoken{}, row.Errf(0, "적힌 꼴에는 한글 음절이 하나는 있어야 한다")
	}
	if lead < 0 {
		lead = rule.LeadVowel
	}
	if !strings.Contains(row.Cells[1], ":") {
		return spoken{}, row.Errf(1, "표준 표지는 색인의 이름(맺음:고, 보조사:도)으로 적는다")
	}
	flags, err := row.Flags(2, spokenFlags)
	if err != nil {
		return spoken{}, err
	}
	return spoken{form, lead, row.Cells[1], flags&fNeedIda != 0}, nil
}
