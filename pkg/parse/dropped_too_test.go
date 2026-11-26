//ff:func feature=parse type=reader control=sequence
//ff:what TestDroppedToo — with a word table a whole tail read after a said stem 이 that has a letter before it is read with the stem dropped too (고양이다); not without a word table, not for a dropped stem, not under a fused piece, and not when the syllable 이 begins the unit (이다)
package parse

import "testing"

func TestDroppedToo(t *testing.T) {
	sc := wordsScan("고양이다", layerCost)
	if !sc.droppedToo(tail{start: 3}, sc.ix.copula) {
		t.Fatal("고양이다 with a word table")
	}
	if sc.droppedToo(tail{start: 3}, sc.ix.dropped) || sc.droppedToo(tail{start: 3, base: '이'}, sc.ix.copula) || sc.droppedToo(tail{start: 1}, sc.ix.copula) {
		t.Fatal("a dropped stem, a fused piece, a syllable 이 that begins the unit")
	}
	bare := scanOf("고양이다")
	if bare.droppedToo(tail{start: 3}, bare.ix.copula) {
		t.Fatal("without a word table")
	}
}
