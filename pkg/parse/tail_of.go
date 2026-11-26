//ff:func feature=parse type=reader control=selection
//ff:what tailOf — what kind of tail an entry makes after a stem: a case particle, the auxiliary particle 요, another particle or suffix, a closing only verbs take, a closing verbs and 있다, 없다 and 계시다 take (existClosings), or another piece of a predicate tail
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func tailOf(e *entry) int {
	switch {
	case e.holds:
		return rule.TailCopula
	case e.tail != 0:
		return e.tail
	case slices.Contains(verbClosings, e.name):
		return rule.TailVerb
	case slices.Contains(existClosings, e.name):
		return rule.TailExist
	case e.class != classParticle:
		return rule.TailPred
	case e.kase != 0:
		return rule.TailCase
	case e.name == yoName:
		return rule.TailPolite
	}
	return rule.TailAux
}
