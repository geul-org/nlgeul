//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestSpansOf — the spans of all the ways once each, by start and then end
package token

import "testing"

func TestSpansOf(t *testing.T) {
	a := cut{spans: []Span{{0, 3}, {3, 6}}}
	b := cut{spans: []Span{{0, 1}, {1, 6}}}
	c := cut{spans: []Span{{0, 3}, {3, 6}}}
	got := spansOf([]cut{a, b, c})
	want := []Span{{0, 1}, {0, 3}, {1, 6}, {3, 6}}
	if len(got) != len(want) {
		t.Fatal(got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatal(got)
		}
	}
	if len(spansOf(nil)) != 0 {
		t.Fatal("no way, no span")
	}
}
