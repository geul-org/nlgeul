//ff:func feature=token type=reader control=iteration dimension=1
//ff:what afterClosings — the tails that end with an auxiliary particle right after a linking closing (보+지+도, 먹+고+는): each predicate tail that ends before the particle with such a closing, judged by the graph of the pair, with the particle as one more piece; not for a pair that is a closing of the tables itself (어도, 지만), nor for a particle after another particle
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
)

func (sc *scan) afterClosings(start, depth int, pc piece, entries []*entry) []tail {
	name := auxName(entries)
	if name == "" || depth != 0 || start < 2 {
		return nil
	}
	pc.names = name
	var out []tail
	for _, t := range sc.closings(start, false) {
		last := t.pieces[len(t.pieces)-1]
		pair := strings.NewReplacer("(", "", ")", "").Replace(last.text) + pc.text
		if sc.ix.closingSet[pair] || sc.judge(sc.ix.linkAux, start, start+len([]rune(pc.text)), rule.Piece{Form: last.names, After: name, HasStem: true, Next: -1}) <= 0 {
			continue
		}
		out = append(out, joined([]tail{t}, pc)...)
	}
	return out
}
