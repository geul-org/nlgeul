//ff:func feature=token type=reader control=sequence
//ff:what TestBefore — without a word table: by verdict, then by preference; with one: by preference plus the log of the verdict
package token

import "testing"

func TestBefore(t *testing.T) {
	none := scanOf("x").ix
	hi, lo := &path{verdict: 1, pref: -9}, &path{verdict: 0.3, pref: -1}
	if !none.before(hi, lo) || none.before(lo, hi) || !none.before(&path{verdict: 1, pref: -1}, hi) {
		t.Fatal("without a table")
	}
	ix := wordsScan("x", layerCost).ix
	if !ix.before(lo, hi) || !ix.before(&path{verdict: 1, pref: -5}, &path{verdict: 0.1, pref: -4.5}) {
		t.Fatal("with a table")
	}
}
