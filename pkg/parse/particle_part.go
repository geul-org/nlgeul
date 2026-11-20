//ff:func feature=parse type=helper control=sequence
//ff:what particlePart — is a part a particle (a case particle, the vocative or an auxiliary particle); the copula is not one
package parse

import "slices"

func particlePart(p Part) bool {
	return slices.Contains(particleKinds, p.Name.Kind)
}
