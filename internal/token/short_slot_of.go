//ff:func feature=token type=reader control=sequence
//ff:what shortSlotOf — one quotation shortened with the past fused in read from a line of its table: its form, what must stand before it, the form it is short for and the stems it follows; it stands where the past stands and is named with its form and the form it is short for
package token

import (
	"github.com/geul-org/nlgeul/data"
	"github.com/geul-org/nlgeul/internal/parse"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func shortSlotOf(row data.Row) (entry, error) {
	lead, form, err := parse.FormOf(row, 0)
	if err != nil {
		return entry{}, err
	}
	if lead < 0 || form == "" {
		return entry{}, row.Errf(0, "선어말의 꼴은 한글 음절이나 받침 뒤에 한글 음절로 적는다")
	}
	after, err := row.Pick(1, parse.AfterNames)
	if err != nil {
		return entry{}, err
	}
	tail, err := row.Pick(3, tailNames)
	if err != nil {
		return entry{}, err
	}
	name := slotLabel + row.Cells[0] + "(" + row.Cells[2] + ")"
	return entry{name: name, form: form, lead: lead, after: after, slot: rule.SlotPast, tail: tail}, nil
}
