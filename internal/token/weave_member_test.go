//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveMember — a line of spans is a way when it begins where the portion begins, ends where it ends, every span is one where states stand and the states chain; a line whose spans fit but whose states do not chain, one with an unknown span, one that begins or ends elsewhere, and an empty one are not
package token

import "testing"

func TestWeaveMember(t *testing.T) {
	w, _ := oddWeave()
	if !w.member([]Span{{0, 1}, {1, 2}, {2, 3}, {3, 4}}) || !w.member([]Span{{0, 2}, {2, 4}}) {
		t.Fatal("the ways")
	}
	if w.member([]Span{{0, 2}, {2, 3}, {3, 4}}) {
		t.Fatal("the states do not chain")
	}
	if w.member([]Span{{0, 2}, {1, 4}}) || w.member([]Span{{0, 9}, {9, 4}}) {
		t.Fatal("an unknown span")
	}
	if w.member([]Span{{1, 2}, {2, 4}}) || w.member([]Span{{0, 2}, {2, 3}}) || w.member(nil) {
		t.Fatal("begins or ends elsewhere")
	}
}
