//ff:func feature=token type=reader control=sequence
//ff:what fitsAux — article 47: the stem of the word after the boundary is an auxiliary of the lists and the word before it ends in a closing; or it is an auxiliary made of a dependent noun and the word before it ends in a modifying closing
package token

import "slices"

func fitsAux(g gap) bool {
	r := stemPart(g.right).Text
	return (slices.Contains(auxiliaryWords, r) && lastKind(g.left) == "맺음") || (slices.Contains(boundAux, r) && lastKind(g.left) == "꾸밈")
}
