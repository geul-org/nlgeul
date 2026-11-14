//ff:func feature=token type=reader control=sequence
//ff:what stepBefore — does word a come before word b by the order candidates have (before): without a word table by verdict, then by preference; with one by preference and verdict together
package token

import "math"

func (ix *index) stepBefore(a, b step) bool {
	if ix.layers != layerNone {
		return a.pref+math.Log10(a.verdict) > b.pref+math.Log10(b.verdict)
	}
	if a.verdict != b.verdict {
		return a.verdict > b.verdict
	}
	return a.pref > b.pref
}
