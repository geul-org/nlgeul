//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestTailsAt — 먹어요 ends in the closing 어 with 요 as well as in the particle 요; 먹던요 gets no 요 after 던; a word that does not end in 요 has only its own tails; a full stop closes a tail or a bare stem; 이게 is read as a fused form; with a word table 가세요 has the honorific under 세; 3개 has the unit, and 1월~ the range sign after it
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
	stops := scanOf("했다.그").tailsAt(3)
	var bare, after bool
	for _, x := range stops {
		bare = bare || (len(x.pieces) == 1 && x.start == 2 && x.pieces[0].names == "마침표:.")
		after = after || (len(x.pieces) == 2 && x.start == 1)
	}
	if !bare || !after {
		t.Fatal(stops)
	}
	var whole bool
	for _, c := range scanOf("이게").tailsAt(2) {
		whole = whole || c.shown == "이것"
	}
	if !whole {
		t.Fatal("no fused form")
	}
	var honor bool
	for _, c := range wordsScan("가세요", layerCost).tailsAt(3) {
		honor = honor || (len(c.pieces) == 3 && c.pieces[0].names == honorName && c.pieces[2].text == "요")
	}
	if !honor {
		t.Fatal("no honorific under 세")
	}
	if len(scanOf("3개").tailsAt(2)) == 0 || len(scanOf("1월~3월").tailsAt(3)) == 0 {
		t.Fatal("no unit")
	}
}
