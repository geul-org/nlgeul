//ff:func feature=token type=helper control=sequence
//ff:what TestNearIn — a candidate as the rules of a sentence read it is what nearOf gives with the node's Glued: the same candidate in a node written apart is not glued and in a node written onto the word before it is
package token

import "testing"

func TestNearIn(t *testing.T) {
	br := Branch{Words: []Eojeol{{Parts: []Part{{Text: "한", Name: Name{Kind: "관형사", Form: "한"}, Stem: true}}}}}
	apart := nearIn(Node{Contest: Contest{Branches: []Branch{br}}}, br)
	glued := nearIn(Node{Contest: Contest{Glued: true, Branches: []Branch{br}}}, br)
	if apart != nearOf(br) || apart.Glued || apart.Kind != "관형사:한" {
		t.Fatal(apart)
	}
	if !glued.Glued || glued.Stem != "한" || glued.Kind != apart.Kind {
		t.Fatal(glued)
	}
}
