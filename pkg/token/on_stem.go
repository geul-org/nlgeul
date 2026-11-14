//ff:func feature=token type=reader control=sequence
//ff:what onStem — the particle candidates that end at j as they stand right on a stem, by key (as groupAt gives them): those given, unless the unit ends at j and a sign that ends a sentence closes it, or the particle 요 follows; there the graphs are asked once more, told so, and the candidates that still stand come back (문제야. keeps the vocative 야 and not the auxiliary one; 에 before 요 is told that it stands on its stem, and what the stem ends in)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

// A particle after another particle or after a closing is not asked
// so: 에서야. and 하고 나서야. keep their auxiliary 야. A particle
// before 요 is asked again where it stands, since whether it stands
// right on a stem is told only here (the rest of the chain is read after
// it, to the left).
func (sc *scan) onStem(j, next int, pieces map[string]piece, from map[string]*entry) (map[string]piece, map[string]*entry) {
	closing, yo := sc.closed && j == len(sc.s), next == rule.LinkYo
	if !closing && !yo {
		return pieces, from
	}
	again, froms, _ := sc.groupAt(classParticle, j, next, false, closing, yo)
	return again, froms
}
