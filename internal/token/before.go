//ff:func feature=token type=reader control=sequence
//ff:what before — does candidate a come before candidate b: without a word table by verdict, then by preference; with one by preference and verdict together
package token

import "math"

// With a word table verdicts are degrees, not two levels: a candidate's
// place is its preference plus log10 of its verdict, so that a candidate
// much less well founded comes later, and no candidate is left out for
// its verdict alone. The verdict's log10 is weighed once (a weight of
// 1, written nowhere as a number); nlgeul hypo weighs it hypo.Lambda
// times. Nothing records a measurement for the weight used here.
func (ix *index) before(a, b *path) bool {
	if ix.layers != layerNone {
		return a.pref+math.Log10(a.verdict) > b.pref+math.Log10(b.verdict)
	}
	if a.verdict != b.verdict {
		return a.verdict > b.verdict
	}
	return a.pref > b.pref
}
