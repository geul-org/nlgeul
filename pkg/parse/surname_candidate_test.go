//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestSurnameCandidate — article 48 on: a family name of the list at the unit's head right before a title of the list that ends the unit or takes a particle is one more candidate, 성:<letters>, a nominal with verdict 1 (김군은, 남궁씨가, 김과장); none when the rule is off, the stem does not begin the unit, has markers, a piece fused into it, a wrong form, a word shown for it or a number's name, is no family name of the list, or the title is the head of another word (김양식장에) or there is none (김밥은)
package parse

import "testing"

func TestSurnameCandidate(t *testing.T) {
	sc := spacingScan("김군은", "48")
	got := sc.surnameCandidate(0, tail{start: 1})
	if len(got) != 1 || got[0].stem != "김" || got[0].name != "성:김" || got[0].kind != kindNominal || got[0].v != 1 || got[0].wrote != "" {
		t.Fatal("a family name before a title and a particle", got)
	}
	if got := spacingScan("남궁씨가", "48").surnameCandidate(0, tail{start: 2}); len(got) != 1 || got[0].name != "성:남궁" || got[0].stem != "남궁" {
		t.Fatal("a family name of two letters", got)
	}
	if got := spacingScan("김과장", "48").surnameCandidate(0, tail{start: 1}); len(got) != 1 || got[0].name != "성:김" {
		t.Fatal("a title that ends the unit", got)
	}
	if got := spacingScan("김군은", "48다만").surnameCandidate(0, tail{start: 1}); got != nil {
		t.Fatal("the rule of article 48 is off", got)
	}
	if got := spacingScan("김군은").surnameCandidate(0, tail{start: 1}); got != nil {
		t.Fatal("no rule is on", got)
	}
	if got := spacingScan("갑김군은", "48").surnameCandidate(1, tail{start: 2}); got != nil {
		t.Fatal("a family name that does not begin the unit", got)
	}
	pieced := []piece{{"은", "보조사:은", 1, -1, 0, false, ""}}
	for _, tl := range []tail{{start: 1, pieces: pieced}, {start: 1, as: numName}, {start: 1, wrote: "긴"}, {start: 1, shown: "김"}, {start: 1, base: '기'}} {
		if got := sc.surnameCandidate(0, tl); got != nil {
			t.Fatal("a stem with a marker, a number's name, a wrong form, a word shown or a fused piece", tl, got)
		}
	}
	if got := spacingScan("갑군은", "48").surnameCandidate(0, tail{start: 1}); got != nil {
		t.Fatal("a letter that is no family name of the list", got)
	}
	if got := spacingScan("김양식장에", "48").surnameCandidate(0, tail{start: 1}); got != nil {
		t.Fatal("a title that is the head of another word", got)
	}
	if got := spacingScan("김밥은", "48").surnameCandidate(0, tail{start: 1}); got != nil {
		t.Fatal("no title after the family name", got)
	}
}
