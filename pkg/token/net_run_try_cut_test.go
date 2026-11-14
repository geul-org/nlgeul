//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunTryCut — no cut is tried with the window off, at the unit's end, or before windowSize tokens have been worked out from the last cut; later the words every way agrees on are cut off the front of the net, the stretch is the first word, and the places before it are dropped
package token

import "testing"

func TestNetRunTryCut(t *testing.T) {
	r, _ := windowSteps(t, 5, 5)
	if _, ok := r.tryCut(); ok {
		t.Fatal("too early")
	}
	r.window = false
	for r.j < 30 {
		r.step()
	}
	if _, ok := r.tryCut(); ok {
		t.Fatal("window off")
	}
	r.window, r.j = true, r.n
	if _, ok := r.tryCut(); ok {
		t.Fatal("the unit's end")
	}
	r, st := windowCut(t, 5)
	if st.contest.Text != "고양이가" || r.c != 4 || r.best[0] != nil || r.sc.floor != 4 || st.redo == nil || r.stats.cuts != 1 {
		t.Fatal(st.contest.Text, r.c)
	}
}
