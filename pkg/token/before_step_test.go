//ff:func feature=token type=reader control=sequence
//ff:what TestStepBefore — a word comes before another without a word table by verdict, then by preference; with one by preference plus the log of the verdict; equal words come before neither
package token

import "testing"

func TestStepBefore(t *testing.T) {
	none := scanOf("x").ix
	hi, lo := step{verdict: 1, pref: -9}, step{verdict: 0.3, pref: -1}
	if !none.stepBefore(hi, lo) || none.stepBefore(lo, hi) || !none.stepBefore(step{verdict: 1, pref: -1}, hi) || none.stepBefore(hi, hi) {
		t.Fatal("without a table")
	}
	ix := wordsScan("x", layerCost).ix
	if !ix.stepBefore(lo, hi) || ix.stepBefore(hi, lo) || ix.stepBefore(hi, hi) {
		t.Fatal("with a table")
	}
}
