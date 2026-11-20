//ff:func feature=parse type=reader control=sequence
//ff:what TestPrepare — every run is an atom, joined runs one atom; a number with 제 before it begins at 제 and its end is noted (제3조│제1항); a unit of Hangul has none; the places where a word ends in a modifying closing, a determiner or the linking 어/아 start empty
package parse

import (
	"reflect"
	"testing"
)

func TestPrepare(t *testing.T) {
	sc := scanOf("v1.2.3이고 L2는")
	sc.prepare()
	if len(sc.atoms) != 2 || !reflect.DeepEqual(sc.atoms[0].names, []string{"버전"}) || sc.atoms[0].end != 6 || sc.atoms[9].end != 11 {
		t.Fatal(sc.atoms)
	}
	sc = scanOf("제3조제1항")
	sc.prepare()
	if len(sc.atoms) != 2 || sc.atoms[0].end != 2 || sc.atoms[3].end != 5 || !sc.ordEnds[2] || !sc.ordEnds[5] || len(sc.ordEnds) != 2 {
		t.Fatal(sc.atoms, sc.ordEnds)
	}
	sc = scanOf("한글만")
	sc.prepare()
	if len(sc.atoms) != 0 || len(sc.inside) != 4 {
		t.Fatal(sc.atoms)
	}
	if sc.adnEnds == nil || sc.detEnds == nil || sc.connEnds == nil || len(sc.adnEnds)+len(sc.detEnds)+len(sc.connEnds) != 0 {
		t.Fatal("ends before reading", sc.adnEnds, sc.detEnds, sc.connEnds)
	}
}
