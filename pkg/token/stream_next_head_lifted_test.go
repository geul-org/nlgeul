//ff:func feature=token type=reader control=sequence
//ff:what TestStreamNextHeadLifted — NextHead gives the first head candidates after the lifts: with complementRule the one candidate of 학생이 before 되었다 is the raised 보격 이, two are 보격 then 주격, and they are the first of what Next gives
package token

import "testing"

func TestStreamNextHeadLifted(t *testing.T) {
	full, _ := sentenceStream(t, "학생이 되었다", false, complementRule()).Next()
	one, _ := sentenceStream(t, "학생이 되었다", false, complementRule()).NextHead(1)
	two, _ := sentenceStream(t, "학생이 되었다", false, complementRule()).NextHead(2)
	if bs := one.Nodes[0].Contest.Branches; len(bs) != 1 || bs[0].String() != "학생[열쇠:체언]+이[보격:이]" || bs[0].Pref != full.Nodes[0].Contest.Branches[0].Pref {
		t.Fatal(bs)
	}
	if bs := two.Nodes[0].Contest.Branches; len(bs) != 2 || bs[1].String() != "학생[열쇠:체언]+이[주격:이]" || len(full.Nodes[0].Contest.Branches) != 3 || one.Lambda != full.Lambda {
		t.Fatal(bs)
	}
}
