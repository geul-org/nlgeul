//ff:func feature=token type=helper control=sequence
//ff:what TestNodesOf — a node for each span: its letters cut out of the region's, the unit of the region, the candidates of the span best first by lambda, and Glued as the region's for the node that begins the region and true for the others
package token

import "testing"

func TestNodesOf(t *testing.T) {
	a := streamCut([]Span{{4, 6}, {6, 9}}, []string{"가나", "다라마"}, []float64{-1, -2})
	b := streamCut([]Span{{4, 6}, {6, 9}}, []string{"가나", "다라마"}, []float64{-3, -1})
	b.words[0].Words[0].Parts[0].Text = "가나 "
	head := Contest{Unit: 2, Glued: false, Text: "가나다라마"}
	nodes := nodesOf(head, 4, spansOf([]cut{a, b}), poolOf([]cut{a, b}))
	if len(nodes) != 2 || nodes[0].Span != (Span{4, 6}) || nodes[1].Span != (Span{6, 9}) {
		t.Fatal(nodes)
	}
	if c := nodes[0].Contest; c.Text != "가나" || c.Unit != 2 || c.Glued || len(c.Branches) != 2 || c.Branches[0].Pref != -1 {
		t.Fatal(c)
	}
	if c := nodes[1].Contest; c.Text != "다라마" || !c.Glued || len(c.Branches) != 1 {
		t.Fatal(c)
	}
	glued := nodesOf(Contest{Glued: true, Text: "가나다라마"}, 4, spansOf([]cut{a}), poolOf([]cut{a}))
	if !glued[0].Contest.Glued {
		t.Fatal("the first node keeps Glued of the region")
	}
}
