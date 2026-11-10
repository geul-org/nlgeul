//ff:func feature=token type=reader control=selection
//ff:what tailOf — what kind of tail an entry makes after a stem: a case particle, another particle or suffix, or a piece of a predicate tail
package token

import "github.com/geul-org/koparse/pkg/rule"

func tailOf(e *entry) int {
	switch {
	case e.class != classParticle:
		return rule.TailPred
	case e.kase != 0:
		return rule.TailCase
	}
	return rule.TailAux
}
