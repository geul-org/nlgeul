//ff:func feature=token type=reader control=sequence
//ff:what rank — a reading's place among the readings of its unit as one number: its preference plus log10 of its verdict
package token

import "math"

func rank(p *path) float64 {
	return p.pref + math.Log10(p.verdict)
}
