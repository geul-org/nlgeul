//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what afterNaming — with a word table, a copula tail read once more after each predicate tail that ends where the tail starts with a naming closing of copulaNamings (음, ㅁ: 없+음+이+라, 가리키+(ㅁ)+이+(ㄴ), 기다리+(었)+음+이+(ㄹ)): not after a stem of one syllable (함이다, 됨이다: the tail must start at the third letter or later); the word a naming closing makes is a nominal, so the copula follows it as a particle does; the copula's pieces make a predicate tail there, for the stem before the closing is a predicate's; the tail keeps what it tells of its closing, and the gate is the naming closing's, not the one a dropped stem asks of a nominal before it
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) afterNaming(t tail) []tail {
	if !sc.ix.lexed || t.base != 0 || t.start < 2 || !copulaNamingEnd(sc.s[t.start-1]) {
		return nil
	}
	rest := append([]piece{}, t.pieces...)
	for i := range rest {
		rest[i].tail = rule.TailPred
	}
	var out []tail
	for _, l := range sc.closings(t.start, false, false) {
		if !slices.Contains(copulaNamings, l.pieces[len(l.pieces)-1].names) {
			continue
		}
		l = joined([]tail{l}, rest...)[0]
		l.noYo, l.close = t.noYo, t.close
		out = append(out, l)
	}
	return out
}
