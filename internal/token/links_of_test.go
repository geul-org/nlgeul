//ff:func feature=token type=helper control=sequence
//ff:what TestLinksOf — for each node the places of the nodes after it, the nearest first whatever order the pairs came in
package token

import "testing"

func TestLinksOf(t *testing.T) {
	nodes := []Node{{Span: Span{0, 1}}, {Span: Span{1, 2}}, {Span: Span{1, 3}}, {Span: Span{2, 4}}, {Span: Span{3, 4}}}
	edges := []edge{{Span{0, 1}, Span{1, 3}}, {Span{0, 1}, Span{1, 2}}, {Span{1, 2}, Span{2, 4}}, {Span{1, 3}, Span{3, 4}}}
	got := linksOf(nodes, edges)
	if len(got) != 5 || len(got[0]) != 2 || got[0][0] != 1 || got[0][1] != 2 || len(got[1]) != 1 || got[1][0] != 3 || len(got[2]) != 1 || got[2][0] != 4 || len(got[3])+len(got[4]) != 0 {
		t.Fatal(got)
	}
}
