//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetRunCommit — a cut is made at a place: how wide the window was is noted, the cut is where the next stretch begins, the stretch is given what it needs to be read again (not the first stretch of letters after a sign of a pair, which grows from planted words, but the ones after it are), and the net is kept from the cut on
package token

import "testing"

func TestNetRunCommit(t *testing.T) {
	r, st := windowCut(t, 5)
	if r.c != 4 || len(r.stats.widths) != 1 || r.stats.widths[0] != 24 || st.redo == nil || st.redo.from != 0 || st.redo.to != 4 || len(st.redo.seeds) != 0 || !st.redo.first {
		t.Fatal(r.c, r.stats.widths, st.redo)
	}
	planted, _ := windowRun(t, "words-far.tsv", windowText(5))
	planted.plant = func(j int) []*path { return nil }
	for k := 0; k < 2; {
		st, ok := planted.tryCut()
		if ok && (k == 0) != (st.redo == nil) {
			t.Fatal("the redo of the cut", k, st.redo)
		}
		if ok {
			k++
		} else {
			planted.step()
		}
	}
}
