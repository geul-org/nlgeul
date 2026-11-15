//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestBoundAfterAdn — the dependent noun 것 or its spoken form 거 where some candidate has a word that ends in a modifying closing end right before it is one written onto its modifier; not at the unit's head, not where no such word ends, and not for another dependent noun (수, 데) or another stem
package token

import "testing"

func TestBoundAfterAdn(t *testing.T) {
	sc := scanOf("하는거")
	if sc.boundAfterAdn(2, "거") {
		t.Fatal("no word that ends in a modifying closing ends before it")
	}
	sc.adnEnds[0], sc.adnEnds[2] = true, true
	for _, stem := range []string{"거", "것"} {
		if !sc.boundAfterAdn(2, stem) {
			t.Error("right after a word that ends in a modifying closing", stem)
		}
	}
	if sc.boundAfterAdn(0, "거") {
		t.Fatal("at the unit's head")
	}
	for _, stem := range []string{"수", "데", "거기", "가"} {
		if sc.boundAfterAdn(2, stem) {
			t.Error("not 것 or 거", stem)
		}
	}
}
