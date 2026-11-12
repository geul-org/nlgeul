//ff:func feature=token type=rule control=sequence
//ff:what TestAfterAdnKind — a stem of two kinds right after a word that ends in a modifying closing is turned down by the kind a reading gives it: as the kind that only stands alone it is, as the noun it also is it is not; with no kind given it stands when any of them does
package token

import "testing"

func TestAfterAdnKind(t *testing.T) {
	sc := wordsScan("먹을지금", layerCost)
	sc.ix.lex["지금"] = word{kinds: kindNominal | kindAlone, count: 10, by: [6]int{5, 0, 0, 0, 0, 5}}
	sc.adnEnds[2] = true
	at := tail{start: 4}
	if sc.afterAdn(2, at, false, 0) {
		t.Fatal("it stands as a noun there", 0)
	}
	if sc.afterAdn(2, at, false, kindNominal) {
		t.Fatal("as a noun")
	}
	if !sc.afterAdn(2, at, false, kindAlone) {
		t.Fatal("as the kind that only stands alone it is turned down")
	}
}
