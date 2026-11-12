//ff:func feature=token type=reader control=sequence
//ff:what TestOnStem — the particle readings right on a stem are those given unless the unit ends there before a sign that ends a sentence; there the auxiliary 야 is dropped from the piece and the vocative stays, 이야 is gone altogether, and a place inside the unit is not asked again
package token

import "testing"

func TestOnStem(t *testing.T) {
	sc := scanOf("학교야")
	pieces, from, _ := sc.readGroup(classParticle, 3, -1, false)
	_, merged, kinds := byForm(pieces, from)
	if merged["야"].names != "보조사:야|호격:야" {
		t.Fatal(merged)
	}
	if ends, endKinds := sc.onStem(3, -1, merged, kinds); ends["야"].names != merged["야"].names || len(endKinds["야"]) != 2 {
		t.Fatal("a unit that no sign closes is read as it is", ends)
	}
	sc.closed = true
	ends, endKinds := sc.onStem(3, -1, merged, kinds)
	if ends["야"].names != vocName || len(endKinds["야"]) != 1 || endKinds["야"][0].gate != sc.ix.call {
		t.Fatal(ends, endKinds)
	}
	if merged["야"].names != "보조사:야|호격:야" || len(kinds["야"]) != 2 {
		t.Fatal("the readings given are left as they were")
	}
	inner, innerFrom, _ := sc.readGroup(classParticle, 2, -1, false)
	_, in, inKinds := byForm(inner, innerFrom)
	if got, _ := sc.onStem(2, -1, in, inKinds); len(got) != len(in) {
		t.Fatal("a place inside the unit is not asked again")
	}

	sc = scanOf("마당이야")
	sc.closed = true
	pieces, from, _ = sc.readGroup(classParticle, 4, -1, false)
	_, merged, kinds = byForm(pieces, from)
	if _, ok := merged["이야"]; !ok {
		t.Fatal(merged)
	}
	if ends, _ := sc.onStem(4, -1, merged, kinds); ends["이야"].names != "" || ends["야"].names != vocName {
		t.Fatal("이야 right on a stem before a sign that ends a sentence", ends)
	}
}
