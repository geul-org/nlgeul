//ff:func feature=token type=helper control=sequence
//ff:what TestWeaveNext — the states at a span that have a state of the set before them, in the order made; empty when none has, and a stale set is not confused with the one before
package token

import "testing"

func TestWeaveNext(t *testing.T) {
	w, _ := oddWeave()
	first := w.of[w.at[Span{0, 1}]]
	if got := w.next(first, w.at[Span{1, 2}]); len(got) != 1 || w.node[got[0]] != w.at[Span{1, 2}] {
		t.Fatal(got)
	}
	if got := w.next(first, w.at[Span{2, 3}]); len(got) != 0 {
		t.Fatal("no state at (2,3) follows (0,1)", got)
	}
	if got := w.next(w.of[w.at[Span{0, 2}]], w.at[Span{1, 2}]); len(got) != 0 {
		t.Fatal("a set used before is not used again", got)
	}
}
