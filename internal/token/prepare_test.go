//ff:func feature=token type=reader control=sequence
//ff:what TestPrepare — every run is an atom, joined runs one atom; a unit of Hangul has none
package token

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
	sc = scanOf("한글만")
	sc.prepare()
	if len(sc.atoms) != 0 || len(sc.inside) != 4 {
		t.Fatal(sc.atoms)
	}
}
