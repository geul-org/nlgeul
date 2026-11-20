//ff:func feature=parse type=reader control=sequence
//ff:what TestPlaceSlotAfterSign — a slot right after a closing sign has no stem before it (차(1대)였다: the 였 is no slot there), while a particle right after it has one, and a slot right after a stem has one
package parse

import "testing"

func TestPlaceSlotAfterSign(t *testing.T) {
	slot := &entry{name: "선어말:였", form: "였", class: classSlot}
	if start, p, ok := scanOf("차(1대)였다").place(slot, 6, false); !ok || start != 5 || p.HasStem || p.Prev != ')' {
		t.Fatal("no stem for a slot right after a closing sign", start, p, ok)
	}
	if _, p, ok := scanOf("차(1대)를").place(&entry{form: "를", class: classParticle}, 6, false); !ok || !p.HasStem {
		t.Fatal("a stem before a particle right after a closing sign", p)
	}
	if _, p, ok := scanOf("하였다").place(slot, 2, false); !ok || !p.HasStem {
		t.Fatal("a stem before a slot right after its stem", p)
	}
}
