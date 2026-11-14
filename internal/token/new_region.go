//ff:func feature=token type=helper control=sequence
//ff:what newRegion — a region from the ways of cutting a contest's stretch (cutsOf gives them for a contest; ways from another source may be given here): its nodes, the words that follow each, the best that can be reached from each, the ways there are (whether a line of spans is one, and how many), and the frontier with the words that begin it; head gives the unit, the letters and Glued, and at is where the stretch begins in the sentence
package token

import "math/big"

func newRegion(head Contest, at int, cuts []cut) *region {
	p, forms := poolOf(cuts), formsOf(cuts)
	nodes := nodesOf(head, at, spansOf(cuts), p)
	best := bestOf(nodes)
	next := linksOf(nodes, pairsOf(cuts))
	rg := &region{
		span:    Span{Start: at, End: at + len([]rune(head.Text))},
		nodes:   nodes,
		best:    best,
		rest:    restOf(best, next),
		next:    next,
		member:  memberOf(forms),
		ways:    big.NewInt(int64(len(forms))),
		clashes: p.clashes,
	}
	rg.start()
	return rg
}
