//ff:func feature=token type=reader control=sequence
//ff:what TestGated — a tail with no graph passes; one that asks for a predicate's stem passes after a verb of the table and not after a noun or a word the table does not have
package token

import "testing"

func TestGated(t *testing.T) {
	sc := wordsScan("가신 학교신 곰신", layerCost)
	if !sc.gated(0, tail{start: 1}) {
		t.Fatal("no gate")
	}
	if !sc.gated(0, tail{start: 1, gate: sc.ix.pred}) || sc.gated(3, tail{start: 5, gate: sc.ix.pred}) || sc.gated(7, tail{start: 8, gate: sc.ix.pred}) {
		t.Fatal("pred")
	}
	if !sc.gated(3, tail{start: 5, gate: sc.ix.nominal}) {
		t.Fatal("nominal")
	}
}
