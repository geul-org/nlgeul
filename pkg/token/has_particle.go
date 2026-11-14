//ff:func feature=token type=helper control=iteration dimension=1
//ff:what hasParticle — does a word have a particle or the copula among the parts after its stem
package token

func hasParticle(e Eojeol) bool {
	for _, p := range tailParts(e) {
		if particlePart(p) || p.Name.Kind == "서술격" {
			return true
		}
	}
	return false
}
