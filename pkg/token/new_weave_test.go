//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestNewWeave — a portion laid out: the spans where states stand sorted by start and then end, the place of each span, which states stand at each (in the order made), the span of each state, and the states before each that lie in the portion
package token

import "testing"

func TestNewWeave(t *testing.T) {
	w, s := oddWeave()
	want := []Span{{0, 1}, {0, 2}, {1, 2}, {2, 3}, {2, 4}, {3, 4}}
	if len(w.spans) != 6 || w.from != 0 || w.to != 4 || len(w.states) != 6 || len(w.mark) != 6 {
		t.Fatal(w.spans)
	}
	for i, sp := range want {
		if w.spans[i] != sp || w.at[sp] != i || len(w.of[i]) != 1 {
			t.Fatal(i, w.spans[i], sp)
		}
	}
	if w.states[w.of[3][0]] != s["y"] || w.node[w.of[3][0]] != 3 || len(w.preds[w.of[3][0]]) != 1 || w.states[w.preds[w.of[3][0]][0]] != s["z"] {
		t.Fatal("y", w.of, w.preds)
	}
}
