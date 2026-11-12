//ff:func feature=token type=reader control=sequence
//ff:what TestExpandMid — the branches of every bundle are read by the candidates of the signs between letters, and the bundles given are not changed; a unit with no such sign gives its bundles back
package token

import "testing"

func TestExpandMid(t *testing.T) {
	sc := scanOf("가·나")
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "·", Name: Name{Kind: midName}}}, End: 2}}, Verdict: 1, Pref: -3}
	in := []Bundle{{Text: "가·나", Branches: []Branch{br}}}
	got := sc.expandMid(in)
	n := got[0].Branches[0].Words[0].Parts[0].Name
	if len(got) != 1 || len(got[0].Branches) != 1 || n.Kind != "가운뎃점" || got[0].Branches[0].Pref != -3 {
		t.Fatal(got)
	}
	if in[0].Branches[0].Words[0].Parts[0].Name.Kind != midName {
		t.Fatal("changed", in)
	}
	none := scanOf("가나")
	if got = none.expandMid(in); got[0].Branches[0].Words[0].Parts[0].Name.Kind != midName {
		t.Fatal("no sign", got)
	}
}
