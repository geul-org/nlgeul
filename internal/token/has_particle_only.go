//ff:func feature=token type=helper control=iteration dimension=1
//ff:what hasParticleOnly — does a word have a particle (the copula not counted) among the parts after its stem
package token

func hasParticleOnly(e Eojeol) bool {
	for _, p := range tailParts(e) {
		if particlePart(p) {
			return true
		}
	}
	return false
}
