//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ranked — whether the segments have ranks 1, 2, 3 … in order and their lambdas never rise (to a hair of float rounding)
package token

func ranked(segs []Segment) bool {
	for i, sg := range segs {
		if sg.Rank != i+1 || (i > 0 && sg.Lambda > segs[i-1].Lambda+1e-9) {
			return false
		}
	}
	return true
}
