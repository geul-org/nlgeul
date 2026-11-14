//ff:func feature=token type=reader control=sequence
//ff:what TestCommonSpans — the spans every line begins with, in order; none for no line, and for lines that differ at the first span
package token

import "testing"

func TestCommonSpans(t *testing.T) {
	a := []Span{{0, 2}, {2, 5}, {5, 6}}
	b := []Span{{0, 2}, {2, 5}}
	c := []Span{{0, 2}, {2, 4}, {4, 6}}
	if got := commonSpans(nil); got != nil {
		t.Fatal(got)
	}
	if got := commonSpans([][]Span{a}); len(got) != 3 {
		t.Fatal(got)
	}
	if got := commonSpans([][]Span{a, b}); len(got) != 2 {
		t.Fatal(got)
	}
	if got := commonSpans([][]Span{a, c}); len(got) != 1 || got[0] != (Span{0, 2}) {
		t.Fatal(got)
	}
	if got := commonSpans([][]Span{{{0, 3}}, {{0, 2}}}); len(got) != 0 {
		t.Fatal(got)
	}
}
