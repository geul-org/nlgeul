//ff:func feature=token type=helper control=iteration dimension=1
//ff:what windowCut — a net run over windowText(times), worked out until the window cuts a stretch off, with the stretch
package token

import "testing"

func windowCut(t *testing.T, times int) (*netRun, unitStretch) {
	t.Helper()
	r, _ := windowRun(t, "words-far.tsv", windowText(times))
	for {
		if st, ok := r.tryCut(); ok {
			return r, st
		}
		if r.j == r.n {
			t.Fatal("no cut")
		}
		r.step()
	}
}
