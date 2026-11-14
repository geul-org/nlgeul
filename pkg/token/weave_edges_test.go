//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestWeaveEdges — the pairs of spans where a state at the second follows a state at the first, in the order of the spans
package token

import "testing"

func TestWeaveEdges(t *testing.T) {
	w, _ := oddWeave()
	got := w.edges()
	want := []edge{{Span{0, 1}, Span{1, 2}}, {Span{0, 2}, Span{2, 4}}, {Span{1, 2}, Span{2, 3}}, {Span{2, 3}, Span{3, 4}}}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(i, got[i], want[i])
		}
	}
}
