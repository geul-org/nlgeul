//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunOpen — a place is part of the window's end when it lies less than windowSize tokens from the place worked out, or a long key of the tables still fits the letters from it; the places farther back are not
package token

import "testing"

func TestNetRunOpen(t *testing.T) {
	r, _ := windowSteps(t, 5, 24)
	got := r.open()
	if len(got) != 25 || got[0] || !got[24] || !got[15] || got[14] {
		t.Fatal(got)
	}
	long, _ := windowRun(t, "words-far.tsv", "대한민국임시정부수립기념일을마신다")
	for long.j < 12 {
		long.step()
	}
	if got = long.open(); !got[0] || !got[12] {
		t.Fatal("a long key is open", got)
	}
}
