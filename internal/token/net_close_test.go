//ff:func feature=token type=reader control=iteration dimension=1
//ff:what TestNetClose — the net of a unit closed: the markers written apart from the unit before are among its ends only when no cut was made (from is 0), the states are those kept, and the net knows where its letters begin
package token

import "testing"

func TestNetClose(t *testing.T) {
	r, sc := windowRun(t, "words-far.tsv", "에서")
	r.window = false
	for r.j < r.n {
		r.step()
	}
	whole := sc.netClose(r.best, 0)
	cut := sc.netClose(r.best, 1)
	if whole.from != 0 || cut.from != 1 || whole.n != 2 || len(whole.ends) <= len(cut.ends) || len(cut.all) != len(sc.registry().all) {
		t.Fatal(len(whole.ends), len(cut.ends))
	}
}
