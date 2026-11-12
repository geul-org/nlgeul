//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterClosings — the tails that end with a particle right after a closing: an auxiliary particle after a linking closing (보+지+도, 먹+고+는), not after another particle, with a word table any particle after the naming closing 기 (있+기+는, 받+기+에+는), and a particle after a closing that asks, inside a chain too (맞+는지+를, 가+(ㄹ)지+도); each predicate tail that ends before the particle with such a closing, judged by the graph of the pair, with the particle as one more piece; not for a pair that is a closing of the tables itself (어도, 지만); the tails before a particle after 기 are looked up as chained, so that a copula stem dropped before 기 is read there and nowhere else (문제+(이)+기+도)
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) afterClosings(start, depth int, pc piece, entries []*entry) []tail {
	name := auxName(entries)
	noun := sc.ix.nounAux != nil && start >= 2 && namingEnd(sc.s[start-1])
	ask := start >= 2 && askEnd(sc.s[start-1])
	if start < 2 || (!noun && !ask && (name == "" || depth != 0)) {
		return nil
	}
	aux, end := pc, start+len([]rune(pc.text))
	aux.names = name
	var out []tail
	for _, t := range sc.closings(start, false, noun) {
		last := t.pieces[len(t.pieces)-1]
		pair := strings.NewReplacer("(", "", ")", "").Replace(last.text) + pc.text
		at := rule.Piece{Form: pc.text, Before: last.names, HasStem: true, Next: -1}
		switch {
		case sc.ix.closingSet[pair]:
		case noun && sc.judge(sc.ix.nounAux, start, end, at) > 0:
			out = append(out, joined([]tail{t}, pc)...)
		case ask && sc.judge(sc.ix.askAux, start, end, at) > 0:
			out = append(out, joined([]tail{t}, pc)...)
		case name != "" && depth == 0 && sc.judge(sc.ix.linkAux, start, end, at) > 0:
			out = append(out, joined([]tail{t}, aux)...)
		}
	}
	return out
}
