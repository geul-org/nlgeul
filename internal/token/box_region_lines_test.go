//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestBoxRegionLines — a segment with no node is its letters on one line; one way (Ways 1) is its boxes stacked with nothing at the right; more ways are the boxes with the brackets
package token

import (
	"strings"
	"testing"
)

func TestBoxRegionLines(t *testing.T) {
	letters := []rune("가나다라마")
	none := boxRegionLines([]Segment{{Region: Span{1, 4}}}, letters, 0)
	if len(none) != 1 || none[0] != "나다라" {
		t.Fatal(none)
	}
	one := boxRegionLines([]Segment{boxSegment(1, 1, 0, Span{0, 2}, boxNode(0, 2, "가나", "가나"))}, letters, 0)
	if len(one) != 3 {
		t.Fatal(one)
	}
	for _, l := range one {
		if r, ok := edgeAt(l, cellWidth(l)); !ok || strings.ContainsRune(l, '├') || (r != '┐' && r != '│' && r != '┘') {
			t.Fatalf("%q", l)
		}
	}
	two := boxRegionLines([]Segment{
		boxSegment(1, 2, 0, Span{0, 2}, boxNode(0, 2, "가나", "가나")),
		boxSegment(2, 2, 0, Span{0, 2}, boxNode(0, 1, "가", "가"), boxNode(1, 2, "나", "나")),
	}, letters, 0)
	if len(two) != 3+3+3 || two[0][len(two[0])-len(" ─┐ ① λ 합 0.00"):] != " ─┐ ① λ 합 0.00" {
		t.Fatalf("%q", two)
	}
}
