//ff:func feature=token type=helper control=sequence
//ff:what TestNearIn — a candidate as the rules of a sentence read it is what nearOf gives with the node's Glued and the facts of all its candidates: the same candidate in a node written apart is not glued and in a node written onto the word before it is; a determiner alone is an adnominal by each candidate and may stand before no other adnominal, and beside an adverb's reading it is neither, while an adverb alone takes no adnominal
package token

import "testing"

func TestNearIn(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "한", Name: Name{Kind: "관형사", Form: "한"}, Stem: true}}}}}
	apart := nearIn(Node{Contest: Contest{Branches: []Branch{br}}}, br)
	glued := nearIn(Node{Contest: Contest{Glued: true, Branches: []Branch{br}}}, br)
	want := nearOf(br)
	want.AdnominalAll = true
	if apart != want || apart.Glued || apart.Kind != "관형사:한" || apart.Unmodifiable {
		t.Fatal(apart)
	}
	if !glued.Glued || glued.Stem != "한" || glued.Kind != apart.Kind {
		t.Fatal(glued)
	}
	adverb := Branch{Words: []Eojeol{{Parts: []Part{{Text: "한", Name: Name{Kind: "열쇠", Form: "부사"}, Stem: true}}}}}
	both := nearIn(Node{Contest: Contest{Branches: []Branch{br, adverb}}}, br)
	if both.AdnominalAll || both.Unmodifiable {
		t.Fatal("beside an adverb's reading", both)
	}
	alone := nearIn(Node{Contest: Contest{Branches: []Branch{adverb}}}, adverb)
	if alone.AdnominalAll || !alone.Unmodifiable {
		t.Fatal("an adverb alone", alone)
	}
}
