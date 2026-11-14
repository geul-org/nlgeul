//ff:func feature=token type=helper control=iteration dimension=1
//ff:what windowSteps — a net run over windowText(times) with its scan, worked out for k places (no cut is tried)
package token

import "testing"

func windowSteps(t *testing.T, times, k int) (*netRun, *scan) {
	t.Helper()
	r, sc := windowRun(t, "words-far.tsv", windowText(times))
	for i := 0; i < k; i++ {
		r.step()
	}
	return r, sc
}
