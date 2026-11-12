//ff:func feature=token type=helper control=sequence
//ff:what TestHeadRun — the first run takes the sign kept at the unit's head; not a number, not @ or # (signs of their own), not a sign of another kind, not when the run begins elsewhere, and no runs give none
package token

import "testing"

func TestHeadRun(t *testing.T) {
	if got := scanOf("$ab").headRun([][2]int{{1, 3}}); got[0][0] != 0 || got[0][1] != 3 {
		t.Fatal(got)
	}
	if got := scanOf("$40").headRun([][2]int{{1, 3}}); got[0][0] != 1 {
		t.Fatal("a number", got)
	}
	if got := scanOf("@ab").headRun([][2]int{{1, 3}}); got[0][0] != 1 {
		t.Fatal("@ is a sign of its own", got)
	}
	if got := scanOf("#ab").headRun([][2]int{{1, 3}}); got[0][0] != 1 {
		t.Fatal("# is a sign of its own", got)
	}
	if got := scanOf("-ab").headRun([][2]int{{1, 3}}); got[0][0] != 1 {
		t.Fatal("another sign", got)
	}
	if got := scanOf("가@ab").headRun([][2]int{{2, 4}}); got[0][0] != 2 {
		t.Fatal("not at the head", got)
	}
	if got := scanOf("가").headRun(nil); got != nil {
		t.Fatal(got)
	}
}
