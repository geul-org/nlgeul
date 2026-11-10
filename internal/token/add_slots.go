//ff:func feature=token type=reader control=iteration dimension=1
//ff:what addSlots — the slots before a closing (endings.md section 1) with their order and observed counts, form by form: 시, 었/았/였 (also fused as ㅆ, which is counted as the form it is written out as), 겠, 더
package token

import "github.com/geul-org/koparse/pkg/rule"

var slotRows = []entry{
	{name: "선어말:시", form: "시", slot: rule.SlotHonor, count: honorCount},
	{name: "선어말:었", form: "었", slot: rule.SlotPast, count: 3534371},
	{name: "선어말:았", form: "았", slot: rule.SlotPast, count: 843790},
	{name: "선어말:였", form: "였", slot: rule.SlotPast, count: 3260287},
	{name: "선어말:ㅆ", lead: jongSS, slot: rule.SlotPast},
	{name: "선어말:겠", form: "겠", slot: rule.SlotGuess, count: 36377},
	{name: "선어말:더", form: "더", slot: rule.SlotRecall, count: 150060},
}

func (ix *index) addSlots() {
	for i := range slotRows {
		e := slotRows[i]
		e.class = classSlot
		ix.add(&e)
	}
}
