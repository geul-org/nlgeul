//ff:func feature=parse type=reader control=sequence
//ff:what TestGroupBase — the lead, the form and the place or kind, parted by dots; a closing written wrong has the letters it is written with at the end, so it is never filed with the right closing of the same letters; a closing read only after a word that passes a graph has that graph's name at the end (a gated particle has not), so it is never filed with a closing of the same letters read with no such graph; the entry's own name is not in the key
package parse

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func TestGroupBase(t *testing.T) {
	if got := groupBase(&entry{form: "이", slot: 1, name: "주격:이"}); got != "·이·1" {
		t.Fatal(got)
	}
	if got := groupBase(&entry{form: "께", lead: rule.JongL, slot: 0}); got != "ㄹ·께·0" {
		t.Fatal(got)
	}
	right := groupBase(&entry{form: "께", lead: rule.JongL, slot: 0})
	wrong := groupBase(&entry{form: "께", lead: rule.JongL, slot: 0, wrote: "ㄹ께"})
	if wrong != "ㄹ·께·0·오표기:ㄹ께" || wrong == right {
		t.Fatal(wrong, right)
	}
	gate := &entry{name: predGate}
	if got := groupBase(&entry{form: "야", class: classClosing, gate: gate}); got != "·야·0·"+predGate || got == groupBase(&entry{form: "야", class: classClosing}) {
		t.Fatal(got)
	}
	if got := groupBase(&entry{form: "선", class: classParticle, gate: gate}); got != "·선·0" {
		t.Fatal(got)
	}
}
