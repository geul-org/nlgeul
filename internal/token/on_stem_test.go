//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestOnStem — a unit that no sign closes is read as it is (보조사:야 and 호격:야 are two candidates, each its own piece and entry); where a sign that ends a sentence closes the unit, the graphs are asked again and only the vocative 야 stands, with the gate of the vocative (a called word before it); the pieces given are left as they were; a place inside the unit is not asked again; 이야 right on a stem before such a sign is turned down while 야 stands
package token

import "testing"

func TestOnStem(t *testing.T) {
	sc := scanOf("학교야")
	pieces, from, _ := sc.readGroup(classParticle, 3, -1, false)
	if len(pieces) != 2 || pieceNamed(pieces, "보조사:야").names == "" || pieceNamed(pieces, vocName).names == "" {
		t.Fatal(pieces)
	}
	if ends, endFrom := sc.onStem(3, -1, pieces, from); len(ends) != 2 || len(endFrom) != 2 {
		t.Fatal("a unit that no sign closes is read as it is", ends)
	}
	sc.closed = true
	ends, endFrom := sc.onStem(3, -1, pieces, from)
	voc := pieceNamed(ends, vocName)
	if len(ends) != 1 || voc.names != vocName {
		t.Fatal(ends)
	}
	for k, e := range endFrom {
		if e.name != vocName || e.gate != sc.ix.call || ends[k].names != vocName {
			t.Fatal(k, e.name, e.gate)
		}
	}
	if len(pieces) != 2 || len(from) != 2 {
		t.Fatal("the candidates given are left as they were")
	}
	inner, innerFrom, _ := sc.readGroup(classParticle, 2, -1, false)
	if got, _ := sc.onStem(2, -1, inner, innerFrom); len(got) != len(inner) {
		t.Fatal("a place inside the unit is not asked again")
	}

	sc = scanOf("마당이야")
	sc.closed = true
	pieces, from, _ = sc.readGroup(classParticle, 4, -1, false)
	if pieceNamed(pieces, "보조사:이야").names == "" {
		t.Fatal(pieces)
	}
	ends, _ = sc.onStem(4, -1, pieces, from)
	if pieceNamed(ends, "보조사:이야").names != "" || pieceNamed(ends, vocName).names != vocName {
		t.Fatal("이야 right on a stem before a sign that ends a sentence", ends)
	}
}
