//ff:func feature=parse type=reader control=sequence
//ff:what TestJoinShort — 이게 and 된다는게 are 것이 (subject), 이건 것은 (no case), 그걸 것을 (object); not after a plain word, not a word of one syllable, not another ending
package parse

import "testing"

func TestJoinShort(t *testing.T) {
	if c, ok := joinShort([]rune("이게")); !ok || c != (chainFact{4, "게", "게", 0, caseSubj | caseComp, "이+게"}) {
		t.Fatal(c, ok)
	}
	if c, ok := joinShort([]rune("된다는게")); !ok || c.size != 7 || c.marked != "된+다+는+게" {
		t.Fatal(c, ok)
	}
	if c, ok := joinShort([]rune("이건")); !ok || c.kase != 0 || c.marked != "이+건" {
		t.Fatal(c, ok)
	}
	if c, ok := joinShort([]rune("그걸")); !ok || c.kase != caseObj {
		t.Fatal(c, ok)
	}
	if _, ok := joinShort([]rune("물건")); ok {
		t.Fatal("물건")
	}
	if _, ok := joinShort([]rune("게")); ok {
		t.Fatal("게")
	}
	if _, ok := joinShort([]rune("이는")); ok {
		t.Fatal("이는")
	}
}
