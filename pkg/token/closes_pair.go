//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closesPair — may the sign of a pair at a unit's head close a pair that stands open before it: one of its candidates is a closing at a depth of 1 or more (a closing with none of its kind open stands at depth 0)
package token

func closesPair(candidates []pairing) bool {
	for _, c := range candidates {
		if c.name == shutName && c.depth > 0 {
			return true
		}
	}
	return false
}
