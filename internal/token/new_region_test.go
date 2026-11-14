//ff:func feature=token type=helper control=sequence
//ff:what TestNewRegion — a region from two ways of cutting a stretch of five letters: its span, nodes (one for each span), the two ways, the best and the rest for each node, the words that follow each, and a frontier holding the two words that begin it, the better first
package token

import "testing"

func TestNewRegion(t *testing.T) {
	rg := skewRegion()
	if rg.span != (Span{0, 5}) || len(rg.nodes) != 7 || rg.ways.Int64() != 2 || rg.clashes != 0 || rg.sent != 0 || len(rg.out) != 0 {
		t.Fatal(rg)
	}
	if len(rg.best) != 7 || len(rg.rest) != 7 || len(rg.next) != 7 || rg.rest[1] != -2.5 || rg.rest[0] != -3.5 || rg.rest[3] != -1.5 {
		t.Fatal(rg.best, rg.rest, rg.next)
	}
	if len(rg.next[3]) != 2 || rg.next[3][0] != 4 || rg.next[3][1] != 5 {
		t.Fatal(rg.next)
	}
	if len(rg.front.items) != 2 || rg.front.items[0].node != 1 || rg.front.items[1].node != 0 {
		t.Fatal(rg.front.items)
	}
	one := newRegion(Contest{Text: "가나"}, 7, []cut{streamCut([]Span{{7, 9}}, []string{"x"}, []float64{-1})})
	if one.span != (Span{7, 9}) || len(one.nodes) != 1 || len(one.front.items) != 1 {
		t.Fatal(one)
	}
}
