//ff:func feature=token type=helper control=sequence
//ff:what TestRegionSpansAt — the spans of the nodes given by their places
package token

import "testing"

func TestRegionSpansAt(t *testing.T) {
	got := skewRegion().spansAt([]int{1, 3, 5})
	if len(got) != 3 || got[0] != (Span{0, 2}) || got[1] != (Span{2, 3}) || got[2] != (Span{3, 5}) {
		t.Fatal(got)
	}
}
