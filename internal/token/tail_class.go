//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tailClass — what follows a stem in its word, from the pieces after it: a predicate tail with a closing only verbs take, another predicate tail, particles with a case particle, other particles, or nothing
package token

import "github.com/geul-org/koparse/pkg/rule"

func tailClass(pieces []piece) int {
	all := 0
	for _, p := range pieces {
		all |= p.tail
	}
	for _, c := range []int{rule.TailVerb, rule.TailPred, rule.TailCase, rule.TailAux} {
		if all&c != 0 {
			return c
		}
	}
	return rule.TailBare
}
