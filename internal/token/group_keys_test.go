//ff:func feature=token type=helper control=iteration dimension=1
//ff:what TestGroupKeys — the keys of the readings at a place come back by their place or kind, and the readings of one place or kind in the order their entries are in the index (주격 before 보격)
package token

import (
	"reflect"
	"testing"
)

func TestGroupKeys(t *testing.T) {
	sc := scanOf("곰이")
	pieces, from, _ := sc.readGroup(classParticle, 2, -1, false)
	names := []string{}
	for _, k := range sc.groupKeys(pieces, from, 2) {
		names = append(names, pieces[k].names)
	}
	if want := []string{"주격:이", "보격:이"}; !reflect.DeepEqual(names, want) {
		t.Fatal(names)
	}
	apart := map[string]piece{"x": {}, "y": {}}
	by := map[string]*entry{"x": {form: "을", slot: 1}, "y": {form: "가", slot: 1}}
	if got := sc.groupKeys(apart, by, 2); !reflect.DeepEqual(got, []string{"y", "x"}) {
		t.Fatal("places filed apart come in the order of their keys", got)
	}
}
