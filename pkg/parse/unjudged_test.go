//ff:func feature=parse type=helper control=sequence
//ff:what TestUnjudged — the regions not yet judged are those after the last judged one at the end of the queue; with none judged all of them, with the last judged none, and an empty queue gives none
package parse

import "testing"

func TestUnjudged(t *testing.T) {
	j, u1, u2 := &region{judged: true}, &region{}, &region{}
	if got := unjudged([]*region{j, j, u1, u2}); len(got) != 2 || got[0] != u1 || got[1] != u2 {
		t.Fatal(got)
	}
	if got := unjudged([]*region{u1, u2}); len(got) != 2 {
		t.Fatal(got)
	}
	if got := unjudged([]*region{u1, j}); len(got) != 0 {
		t.Fatal(got)
	}
	if got := unjudged(nil); len(got) != 0 {
		t.Fatal(got)
	}
}
