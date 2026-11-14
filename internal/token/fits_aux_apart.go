//ff:func feature=token type=reader control=sequence
//ff:what fitsAuxApart — the proviso of article 47: the stem of the word after the boundary is an auxiliary of the lists and the word before it has a particle after its closing; or the two stems together are an auxiliary made of a dependent noun and the word before it has a particle
package token

import "slices"

func fitsAuxApart(g gap) bool {
	r := stemPart(g.right).Text
	joined := stemPart(g.left).Text + r
	return (slices.Contains(auxiliaryWords, r) && particleAfterClosing(g.left)) || (slices.Contains(boundAux, joined) && hasParticleOnly(g.left))
}
