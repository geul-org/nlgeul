//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTailsAt — 먹어요 ends in the closing 어 with 요 as well as in the particle 요; 먹던요 gets no 요 after 던; a word that does not end in 요 has only its own tails
package token

import "testing"

func TestTailsAt(t *testing.T) {
	sc := scanOf("먹어요")
	var closing, particle bool
	for _, x := range sc.tailsAt(3) {
		closing = closing || (len(x.pieces) == 2 && x.pieces[1].names == "보조사:요" && x.start == 1)
		particle = particle || (len(x.pieces) == 1 && x.start == 2)
	}
	if !closing || !particle {
		t.Fatal(closing, particle)
	}
	sc = scanOf("먹던요")
	for _, x := range sc.tailsAt(3) {
		if len(x.pieces) == 2 {
			t.Fatal(x)
		}
	}
	if got := scanOf("곰이").tailsAt(2); len(got) != 1 {
		t.Fatal(got)
	}
	if got := scanOf("요").tailsAt(1); len(got) != 0 {
		t.Fatal(got)
	}
}
