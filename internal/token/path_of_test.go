//ff:func feature=token type=helper control=sequence
//ff:what TestPathOf — the letters of each node in order; a segment with no node is the letters of its stretch
package token

import (
	"slices"
	"testing"
)

func TestPathOf(t *testing.T) {
	seg := boxSegment(1, 1, 0, Span{0, 6}, boxNode(0, 3, "아버지", "아버지"), boxNode(3, 6, "가방에", "가방에"))
	if got := pathOf(seg, []rune("아버지가방에")); !slices.Equal(got, []string{"아버지", "가방에"}) {
		t.Fatal(got)
	}
	if got := pathOf(Segment{Region: Span{3, 6}}, []rune("아버지가방에")); !slices.Equal(got, []string{"가방에"}) {
		t.Fatal(got)
	}
}
