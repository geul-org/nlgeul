//ff:func feature=token type=helper control=sequence
//ff:what TestFormsOf — the different ways of cutting, one entry for each, by their spans
package token

import "testing"

func TestFormsOf(t *testing.T) {
	a := cut{spans: []Span{{0, 3}, {3, 6}}}
	b := cut{spans: []Span{{0, 3}, {3, 6}}}
	c := cut{spans: []Span{{0, 6}}}
	got := formsOf([]cut{a, b, c})
	if len(got) != 2 || !got["0-3,3-6"] || !got["0-6"] {
		t.Fatal(got)
	}
}
