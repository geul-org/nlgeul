//ff:func feature=parse type=reader control=sequence
//ff:what politeSlotOf — a line of the table of the old polite slots read into an entry: a slot named with its form, which is whole syllables and ends in 오 or 옵, with what must stand before it
package parse

import (
	"github.com/geul-org/nlgeul/pkg/data"
	"github.com/geul-org/nlgeul/pkg/rule"
)

func politeSlotOf(row data.Row) (entry, error) {
	lead, form, err := FormOf(row, 0)
	if err != nil {
		return entry{}, err
	}
	letters := []rune(form)
	if lead != 0 || len(letters) == 0 || (letters[len(letters)-1] != politeSoft && letters[len(letters)-1] != politeHard) {
		return entry{}, row.Errf(0, "공손의 선어말은 오나 옵으로 끝나는 한글 음절로 적는다")
	}
	after, err := row.Pick(1, AfterNames)
	if err != nil {
		return entry{}, err
	}
	return entry{name: slotLabel + form, form: form, class: classSlot, after: after, slot: rule.SlotPolite, wide: true}, nil
}
