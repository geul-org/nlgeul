//ff:func feature=token type=reader control=iteration dimension=1
//ff:what titleAfter — whether a title or an office of the list stands from place j and the unit ends right after it or a particle follows it (right: anything else after it does not count)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) titleAfter(j int) bool {
	for _, w := range titleWords {
		if q := matchAt(sc.s, j, w); q >= 0 && sc.right(q) != rule.RightOther {
			return true
		}
	}
	return false
}
