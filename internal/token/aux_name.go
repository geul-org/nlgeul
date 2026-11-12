//ff:func feature=token type=helper control=sequence
//ff:what auxName — the name of a particle reading when it is one of the auxiliary particles read after a linking closing; "" when it is not
package token

import "slices"

func auxName(e *entry) string {
	if slices.Contains(auxParticles, e.name) {
		return e.name
	}
	return ""
}
