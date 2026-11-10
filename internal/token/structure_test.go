//ff:func feature=token type=reader control=sequence
//ff:what TestStructure — nothing before the third layer; then a cost for each stem and marker, for a boundary inside the unit and for an unseen piece
package token

import "testing"

func TestStructure(t *testing.T) {
	if wordsScan("x", layerKind).ix.structure(2, []piece{{unseen: true}}) != 0 {
		t.Fatal("cost before the third layer")
	}
	ix := wordsScan("x", layerCost).ix
	if got := ix.structure(0, []piece{{}}); got != 2*costPiece {
		t.Fatal(got)
	}
	if got := ix.structure(2, []piece{{unseen: true}}); got != 2*costPiece+costBoundary+costUnseen {
		t.Fatal(got)
	}
}
