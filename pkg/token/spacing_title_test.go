//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestSpacingTitle — article 48: 갑돌씨가 has the candidate 갑돌 씨+가 (갑돌 a new word, 씨 the title of the list) with the new word's value; the two words' preferences add up to the candidate's, and the stream gives the word 씨+가 the same preference
package token

import (
	"math"
	"testing"
)

func TestSpacingTitle(t *testing.T) {
	const line = "갑돌 씨[열쇠:체언]+가[주격:가]"
	r := spacingReader(t, "words-spacing.tsv", "48")
	var found *Branch
	for _, b := range r.Contests("갑돌씨가")[0].Branches {
		if b.String() == line {
			found = &b
		}
	}
	if found == nil || len(found.Words) != 2 || len(found.scores) != 2 || found.Verdict > 0.34 {
		t.Fatal("no candidate", found)
	}
	if math.Abs(found.scores[0].pref+found.scores[1].pref-found.Pref) > 1e-9 {
		t.Fatal("the preferences do not add up", found.scores, found.Pref)
	}
	st := r.Stream("갑돌씨가")
	first, _ := st.Next()
	nodes := nodesAt(streamWays(st, first), Span{2, 4})
	if len(nodes) == 0 {
		t.Fatal("no word 씨+가 in the stream")
	}
	for _, n := range nodes {
		if pref, ok := branchPref(n.Contest.Branches, "씨[열쇠:체언]+가[주격:가]"); ok && math.Abs(pref-found.scores[1].pref) > 1e-9 {
			t.Fatal("the word's preference", pref, found.scores[1].pref)
		}
	}
}
