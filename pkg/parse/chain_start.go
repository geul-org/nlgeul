//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what chainStart — the leftmost start of a particle chain that ends at end, with that chain's particles and the case of its outermost case particle; a chain is at most three particles in the order of their places (particles.md section 5)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// chainStart returns end, 0 and "" when no particle is read there.
// depth is the number of particles already read to the right, right
// the place of the nearest of them. The particles come back joined
// with "+", each followed by one (에서+).
func chainStart(s []rune, end, depth, right int) (int, int, string) {
	best, kase, parts := end, 0, ""
	if depth == 3 {
		return best, kase, parts
	}
	for _, p := range particles {
		start, ok := matchParticle(s, end, p)
		if !ok || !rule.Follows(p.slot, right) {
			continue
		}
		c, inner, left := chainStart(s, start, depth+1, p.slot)
		if c < best {
			best, kase, parts = c, outer(p.kase, inner), left+p.form+"+"
		}
	}
	return best, kase, parts
}
