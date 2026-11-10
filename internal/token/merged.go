//ff:func feature=token type=reader control=sequence
//ff:what merged — two readings of the same token sequence as one piece: both names, the higher verdict and preference
package token

import "math"

func merged(old, add piece) piece {
	if old.names == "" {
		return add
	}
	return piece{add.text, old.names + "|" + add.names, math.Max(old.verdict, add.verdict), math.Max(old.pref, add.pref)}
}
