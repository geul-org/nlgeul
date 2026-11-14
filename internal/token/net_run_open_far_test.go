//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunOpenFar — a place a known word longer than a new word's stem may begin at is part of the window's end when that word ends in the last windowSize tokens; not one before the last cut
package token

import "testing"

func TestNetRunOpenFar(t *testing.T) {
	r, _ := windowRun(t, "words-far.tsv", "대한민국임시정부수립기념일을마신다")
	for r.j < 15 {
		r.step()
	}
	flags := make([]bool, 16)
	r.openFar(flags)
	if !flags[0] {
		t.Fatal("the long word begins at 0", flags)
	}
	r.c = 1
	flags = make([]bool, 16)
	r.openFar(flags)
	if flags[0] {
		t.Fatal("before the last cut")
	}
}
