//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closedBefore — the tails of one closing read once more as an ending right after a final closing of a quoted sentence (closingChains: 있+다+거나, 가+리라+는): each predicate tail that ends where the closing starts, judged by the ending's graph with the name of the tail's last piece, with the closing's pieces after it; the tail keeps its stem's closing, so the stem is judged by the closing before the ending
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) closedBefore(pc piece, e *entry, start, j int) []tail {
	g := sc.ix.chainAux[e.name]
	if g == nil || start < 2 {
		return nil
	}
	rest := partsOf(pc, sc.ix.splitsOf(e.name))
	for i := range rest {
		rest[i].tail = rule.TailPred
	}
	var out []tail
	for _, l := range sc.closings(start, false, false) {
		at := rule.Piece{Form: pc.text, Before: l.pieces[len(l.pieces)-1].names, HasStem: true, Next: -1}
		if sc.judge(g, start, j, at) <= 0 {
			continue
		}
		l = joined([]tail{l}, rest...)[0]
		l.noYo = e.noYo
		out = append(out, l)
	}
	return out
}
