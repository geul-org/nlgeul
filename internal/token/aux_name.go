//ff:func feature=token type=helper control=iteration dimension=1
//ff:what auxName — the name of the reading, among a particle form's readings, that is one of the auxiliary particles read after a linking closing; "" when none is
package token

import "slices"

func auxName(entries []*entry) string {
	for _, e := range entries {
		if slices.Contains(auxParticles, e.name) {
			return e.name
		}
	}
	return ""
}
