//ff:func feature=token type=reader control=sequence
//ff:what fitsParticle — article 41: the word before the boundary has a particle or the copula after its stem
package token

func fitsParticle(g gap) bool {
	return hasParticle(g.left)
}
