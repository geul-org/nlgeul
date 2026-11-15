//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingTitle — article 48: the stream of 갑돌씨가 has the way 갑돌 │ 씨+가 (갑돌 a new word with the new word's value, 씨 the title of the list), the way's lambda is the two words' best lambdas added up, and the word 씨+가 has the preference the stream gives it now (−5.2522)
package token

import (
	"math"
	"testing"
)

func TestSpacingTitle(t *testing.T) {
	r := spacingReader(t, "words-spacing.tsv", "48")
	if !branchLines(r, "갑돌씨가")["갑돌 씨[열쇠:체언]+가[주격:가]"] {
		t.Fatal("no candidate", branchLines(r, "갑돌씨가"))
	}
	st := r.Stream("갑돌씨가")
	first, _ := st.Next()
	ways := streamWays(st, first)
	stem, nodes := nodesAt(ways, Span{0, 2}), nodesAt(ways, Span{2, 4})
	if len(stem) == 0 || stem[0].Contest.Branches[0].String() != "갑돌" || stem[0].Contest.Branches[0].Verdict > 0.34 {
		t.Fatal("no new word 갑돌", stem)
	}
	if len(nodes) == 0 {
		t.Fatal("no word 씨+가 in the stream")
	}
	found := false
	for _, w := range ways {
		if len(w.Nodes) == 2 && w.Nodes[0].Span == (Span{0, 2}) && w.Nodes[1].Span == (Span{2, 4}) {
			found = math.Abs(w.Nodes[0].Contest.Branches[0].Lambda()+w.Nodes[1].Contest.Branches[0].Lambda()-w.Lambda) < 1e-9
		}
	}
	if !found {
		t.Fatal("no way 갑돌 │ 씨+가 whose lambda adds up", ways)
	}
	if pref, ok := branchPref(nodes[0].Contest.Branches, "씨[열쇠:체언]+가[주격:가]"); !ok || math.Abs(pref-(-5.2522)) > 1e-4 {
		t.Fatal("the word's preference", pref, ok)
	}
}
