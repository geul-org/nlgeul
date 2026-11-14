//ff:func feature=token type=reader control=sequence
//ff:what fitsTitle — article 48: the stem of the word after the boundary is a title or an office of the lists
package token

import "slices"

func fitsTitle(g gap) bool {
	return slices.Contains(titleWords, stemPart(g.right).Text)
}
