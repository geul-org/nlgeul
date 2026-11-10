//ff:func feature=token type=reader control=sequence
//ff:what TestReadGroup — the two case readings of 이 after a stem come back as one piece; an entry of another class, one that is not placed and one that is defeated are left out
package token

import "testing"

func TestReadGroup(t *testing.T) {
	sc := scanOf("곰이")
	pieces, from, starts := sc.readGroup(classParticle, 2, -1, false)
	if len(pieces) != 1 || pieces["·이"].names != "주격:이|보격:이" || pieces["·이"].verdict != 1 || starts["·이"] != 1 || from["·이"].slot != 3 {
		t.Fatal(pieces)
	}
	if pieces, _, _ = sc.readGroup(classParticle, 1, -1, false); len(pieces) != 0 {
		t.Fatal(pieces)
	}
	if pieces, _, _ = scanOf("나은").readGroup(classParticle, 2, -1, false); len(pieces) != 0 {
		t.Fatal("은 after a vowel", pieces)
	}
	if pieces, _, _ = sc.readGroup(classClosing, 2, -1, false); len(pieces) != 0 {
		t.Fatal(pieces)
	}
}
