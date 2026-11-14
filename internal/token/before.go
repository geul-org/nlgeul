//ff:func feature=token type=reader control=sequence
//ff:what before — does candidate a come before candidate b: without a word table by verdict, then by preference; with one by preference and verdict together
package token

// With a word table verdicts are degrees, not two levels: a candidate's
// place is its preference plus log10 of its verdict, so that a candidate
// much less well founded comes later, and no candidate is left out for
// its verdict alone. The verdict's log10 is weighed once (a weight of
// 1, written nowhere as a number); internal/combo weighs it combo.Lambda
// times. Nothing records a measurement for the weight used here.
func (ix *index) before(a, b *path) bool {
	return ix.stepBefore(step{verdict: a.verdict, pref: a.pref}, step{verdict: b.verdict, pref: b.pref})
}
