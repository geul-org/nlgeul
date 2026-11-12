//ff:func feature=token type=reader control=sequence
//ff:what TestMeasured — a sign of a pair alone right after a digit gets a branch for each line of the unit table that has it (인치, 초 for ", 피트, 분 for '), sure and with a unit's preference, after the readings it already has; another sign, and no bundle, are given back as they are
package token

import "testing"

func TestMeasured(t *testing.T) {
	ix, _ := newIndex(Options{})
	in := []Bundle{{Text: "\"", Branches: []Branch{{Verdict: 1}}}}
	got := ix.measured(in, unit{text: "\"", head: 1, name: pairName, measure: true})
	if len(got[0].Branches) != 3 || got[0].Branches[0].Words != nil {
		t.Fatal(got)
	}
	b := got[0].Branches[1]
	p := b.Words[0].Parts[0]
	if p.Text != "\"" || p.Stem || len(p.Names) != 1 || p.Names[0].Kind != "단위" || p.Names[0].Form != "인치" || b.Verdict != 1 || b.Pref != ix.afterPrefs[false] || b.Words[0].End != 1 {
		t.Fatal(b)
	}
	if got[0].Branches[2].Words[0].Parts[0].Names[0].Form != "초" {
		t.Fatal("the lines of the table in order", got[0].Branches[2])
	}
	quote := ix.measured([]Bundle{{Text: "'", Branches: []Branch{{}}}}, unit{text: "'", head: 1, name: pairName, measure: true})
	if len(quote[0].Branches) != 3 || quote[0].Branches[1].Words[0].Parts[0].Names[0].Form != "피트" || quote[0].Branches[2].Words[0].Parts[0].Names[0].Form != "분" {
		t.Fatal(quote)
	}
	same := []Bundle{{Text: "\"", Branches: []Branch{{}}}}
	if got = ix.measured(same, unit{text: "\"", head: 1, name: pairName}); len(got[0].Branches) != 1 {
		t.Fatal("not after a digit", got)
	}
	if got = ix.measured(nil, unit{measure: true}); got != nil {
		t.Fatal("no bundle", got)
	}
}
