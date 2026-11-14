//ff:func feature=token type=helper control=sequence
//ff:what netUnit — the scan of the first unit of a sentence over a word table of testdata (none when its name is empty), and its unit and reader
package token

import "testing"

func netUnit(t *testing.T, words, text string) (*scan, unit, *Reader) {
	t.Helper()
	r := streamReader(t, words)
	u := unitsOf(r.ix, text, false)[0]
	return unitScan(r.ix, u), u, r
}
