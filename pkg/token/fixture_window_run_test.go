//ff:func feature=token type=helper control=sequence
//ff:what windowRun — the net run of the first unit of a sentence over a word table of testdata, with its scan and its unit, ready to be worked out (nothing is yet)
package token

import "testing"

func windowRun(t *testing.T, words, text string) (*netRun, *scan) {
	t.Helper()
	sc, u, _ := netUnit(t, words, text)
	return newNetRun(sc, u, true, nil, &netStats{}), sc
}
