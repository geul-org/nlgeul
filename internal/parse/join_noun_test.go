//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinNoun — 아는거 is a bare noun on a modifier; 하는것을 carries its particle; 있는거지 a tail on the noun; nothing when the noun is not there, the left is no modifier, or what follows is neither
package parse

import "testing"

func TestJoinNoun(t *testing.T) {
	if c, tails := joinNoun([]rune("아는거"), 2, "거"); c != (chainFact{4, "", "", 0, caseBare, "아+는+거"}) || tails != nil {
		t.Fatal(c, tails)
	}
	if c, _ := joinNoun([]rune("하는것을"), 2, "것"); c.size != 6 || c.kase != caseObj || c.marked != "하+는+것+을" {
		t.Fatal(c)
	}
	c, tails := joinNoun([]rune("있는거지"), 2, "거")
	if c.size != 0 || len(tails) != 1 || tails[0].size != 4 || tails[0].marked != "있+는+거+지" {
		t.Fatal(c, tails)
	}
	if c, tails := joinNoun([]rune("아는거"), 1, "거"); c.size != 0 || tails != nil {
		t.Fatal("noun not at k")
	}
	if c, _ := joinNoun([]rune("먹고거"), 2, "거"); c.size != 0 {
		t.Fatal("no modifier")
	}
	if c, tails := joinNoun([]rune("할수록"), 1, "수"); c.size != 0 || tails != nil {
		t.Fatal("할수록")
	}
}
