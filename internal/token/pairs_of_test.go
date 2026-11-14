//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestPairsOf — the pairs of words side by side in some way, once each, in the order met
package token

import "testing"

func TestPairsOf(t *testing.T) {
	a := cut{spans: []Span{{0, 1}, {1, 2}, {2, 4}}}
	b := cut{spans: []Span{{0, 1}, {1, 2}, {2, 3}, {3, 4}}}
	c := cut{spans: []Span{{0, 4}}}
	got := pairsOf([]cut{a, b, c})
	want := []edge{{Span{0, 1}, Span{1, 2}}, {Span{1, 2}, Span{2, 4}}, {Span{1, 2}, Span{2, 3}}, {Span{2, 3}, Span{3, 4}}}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(got)
		}
	}
}
