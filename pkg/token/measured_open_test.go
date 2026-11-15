//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestMeasuredOpen — where the sign of a pair may close a pair open before it (one of its candidates a closing at depth 1), the branches of the sign as a unit are added with the verdict of measureGraph lowered (about a third) and the unit's preference; a closing with none of its kind open (depth 0) and an opening leave them sure
package token

import "testing"

func TestMeasuredOpen(t *testing.T) {
	ix, _ := newIndex(Options{})
	read := func(candidates []pairing) Branch {
		got := ix.measured([]Contest{{Text: "'", Branches: []Branch{{Verdict: 1}}}}, unit{text: "'", head: 1, name: pairName, measure: true, candidates: candidates})
		if len(got[0].Branches) != 3 {
			t.Fatal(got)
		}
		return got[0].Branches[1]
	}
	open := read([]pairing{{name: shutName, depth: 1}, {name: openName, depth: 1}})
	if open.Words[0].Parts[0].Name.Form != "피트" || open.Verdict <= 0.3 || open.Verdict >= 0.34 || open.Pref != ix.afterPrefs[false] {
		t.Fatal("a quote open before it", open)
	}
	for _, c := range [][]pairing{{{name: shutName, depth: 0}}, {{name: openName, depth: 1}}, nil} {
		if b := read(c); b.Verdict != 1 {
			t.Error("nothing open", c, b.Verdict)
		}
	}
}
