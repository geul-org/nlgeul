//ff:func feature=token type=helper control=iteration dimension=1
//ff:what lift — add to the preferences of the region's candidates what the rules of a sentence gave them (lifts: by the span of the word and the candidate as printed), put the candidates of each word touched in order by lambda again, and begin the ways of the region anew with the best lambdas as they are now (none is sent yet); the lifts are kept, for a stretch a window cut off is read again with them. Nothing is done with no lifts
package token

func (rg *region) lift(lifts []lift) {
	if len(lifts) == 0 {
		return
	}
	rg.lifts = append(rg.lifts, lifts...)
	for i := range rg.nodes {
		rg.nodes[i] = lifted(rg.nodes[i], lifts)
	}
	rg.best = bestOf(rg.nodes)
	rg.rest = restOf(rg.best, rg.next)
	rg.front, rg.out, rg.sent, rg.skipped = frontier{}, nil, 0, 0
	rg.start()
}
