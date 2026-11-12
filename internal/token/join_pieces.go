//ff:func feature=token type=reader control=sequence
//ff:what joinPieces — two readings of the same token sequence as one piece: both names, the higher verdict and preference (the letters it is written with, if wrong, are kept: only readings of the same letters are joined)
package token

import "math"

func joinPieces(old, add piece) piece {
	if old.names == "" {
		return add
	}
	return piece{add.text, old.names + "|" + add.names, math.Max(old.verdict, add.verdict), math.Max(old.pref, add.pref), old.tail | add.tail, false, add.wrote}
}
