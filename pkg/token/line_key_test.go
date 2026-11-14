//ff:func feature=token type=helper control=sequence
//ff:what TestLineKey — the spans of a way written in order, apart by commas (nothing for no span)
package token

import "testing"

func TestLineKey(t *testing.T) {
	if got := lineKey([]Span{{0, 3}, {3, 6}, {6, 11}}); got != "0-3,3-6,6-11" {
		t.Fatal(got)
	}
	if got := lineKey(nil); got != "" {
		t.Fatal(got)
	}
}
