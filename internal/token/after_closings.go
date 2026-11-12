//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterClosings — the tails that end with a particle right after a closing: an auxiliary particle after a linking closing (보+지+도, 먹+고+는), not after another particle, and with a word table any particle after the naming closing 기 (있+기+는, 받+기+에+는); each predicate tail that ends before the particle with such a closing, judged by the graph of the pair, with the particle as one more piece; not for a pair that is a closing of the tables itself (어도, 지만); while the tails before a particle after 기 are looked up the scan is marked chained, so that a copula stem dropped before 기 is read there and nowhere else (문제+(이)+기+도)
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func (sc *scan) afterClosings(start, depth int, pc piece, entries []*entry) []tail {
	name := auxName(entries)
	noun := sc.ix.nounAux != nil && start >= 2 && sc.s[start-1] == nounLetter
	if start < 2 || (!noun && (name == "" || depth != 0)) {
		return nil
	}
	aux, end := pc, start+len([]rune(pc.text))
	aux.names = name
	var out []tail
	sc.chained = noun
	tails := sc.closings(start, false)
	sc.chained = false
	for _, t := range tails {
		last := t.pieces[len(t.pieces)-1]
		pair := strings.NewReplacer("(", "", ")", "").Replace(last.text) + pc.text
		at := rule.Piece{Form: last.names, After: name, HasStem: true, Next: -1}
		switch {
		case sc.ix.closingSet[pair]:
		case noun && sc.judge(sc.ix.nounAux, start, end, at) > 0:
			out = append(out, joined([]tail{t}, pc)...)
		case name != "" && depth == 0 && sc.judge(sc.ix.linkAux, start, end, at) > 0:
			out = append(out, joined([]tail{t}, aux)...)
		}
	}
	return out
}
