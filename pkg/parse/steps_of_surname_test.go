//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what TestStepsOfSurname — with article 48 on, the family name 김 before 군은 is the word's candidates as they were (a new word without a word table) and then one more, 김[성:김], a nominal with verdict 1; with the rule off it is not there
package parse

import "testing"

func TestStepsOfSurname(t *testing.T) {
	on := spacingScan("김군은", "48").stepsOf(0, 1, tail{start: 1})
	off := spacingScan("김군은").stepsOf(0, 1, tail{start: 1})
	if len(on) != len(off)+1 {
		t.Fatal("one candidate more", len(on), len(off))
	}
	for i, e := range off {
		if on[i].sig != e.sig || on[i].verdict != e.verdict || on[i].pref != e.pref {
			t.Fatal("the candidates before it are kept", on[i].sig, e.sig)
		}
	}
	last := on[len(on)-1]
	if last.sig != "김[성:김]" || last.kind != kindNominal || last.verdict != 1 {
		t.Fatal(last.sig, last.kind, last.verdict)
	}
	for _, e := range off {
		if e.sig == "김[성:김]" {
			t.Fatal("the rule is off", e.sig)
		}
	}
}
