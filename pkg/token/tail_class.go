//ff:func feature=token type=reader control=iteration dimension=1
//ff:what tailClass — what follows a stem in its word, from the pieces after it: the copula and its tail, a predicate tail with a closing only verbs take, one with a closing verbs and 있다, 없다 and 계시다 take (a predicate tail like any other when the slot 시, 었 or 겠 stands before that closing), another predicate tail, particles with a case particle, other particles, or nothing
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func tailClass(pieces []piece) int {
	all, slotted := 0, false
	for _, p := range pieces {
		tail := p.tail
		if tail == rule.TailExist && slotted {
			tail = rule.TailPred
		}
		all |= tail
		slotted = slotted || slices.Contains(existSlots, p.names)
	}
	for _, c := range []int{rule.TailCopula, rule.TailVerb, rule.TailExist, rule.TailPred, rule.TailAdj, rule.TailCase, rule.TailAux} {
		if all&c != 0 {
			return c
		}
	}
	return rule.TailBare
}
