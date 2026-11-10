//ff:func feature=token type=reader control=sequence
//ff:what TestStructure — nothing in the first two layers of a word table; with the third, and with no word table, a cost for the stem (none for the markers), for a boundary inside the unit and for an unseen piece
package token

import "testing"

func TestStructure(t *testing.T) {
	if wordsScan("x", layerKind).ix.structure(2, []piece{{unseen: true}}) != 0 {
		t.Fatal("cost before the third layer")
	}
	ix := wordsScan("x", layerCost).ix
	if got := ix.structure(0, []piece{{}, {}}); got != costPiece {
		t.Fatal(got)
	}
	if got := ix.structure(2, []piece{{unseen: true}}); got != costPiece+costBoundary+costUnseen {
		t.Fatal(got)
	}
	if got := scanOf("x").ix.structure(2, nil); got != costPiece+costBoundary {
		t.Fatal("no word table", got)
	}
}
