//ff:func feature=token type=helper control=iteration dimension=1
//ff:what newRegionNet — a region from a stretch the net read (the same region newRegion makes from the ways of cutting that stretch has): its nodes in the order of their spans, the candidates of each as the cells give them, the words that follow each, the best that can be reached from each, how many ways there are and whether a line of spans is one; head gives the unit, the letters and Glued, and at is where the stretch begins in the sentence
package token

func newRegionNet(head Contest, at int, cs cells) *region {
	delta := at - cs.from
	nodes := make([]Node, 0, len(cs.nodes))
	letters := []rune(head.Text)
	for _, c := range cs.nodes {
		sp := shiftSpans([]Span{c.span}, delta)[0]
		nodes = append(nodes, Node{Span: sp, Contest: Contest{Unit: head.Unit, Glued: sp.Start != at || head.Glued, Text: string(letters[sp.Start-at : sp.End-at]), Branches: c.branches}})
	}
	edges := make([]edge, 0, len(cs.edges))
	for _, e := range cs.edges {
		edges = append(edges, edge{from: shiftSpans([]Span{e.from}, delta)[0], to: shiftSpans([]Span{e.to}, delta)[0]})
	}
	best := bestOf(nodes)
	next := linksOf(nodes, edges)
	rg := &region{
		span:    Span{Start: at, End: at + len([]rune(head.Text))},
		nodes:   nodes,
		best:    best,
		rest:    restOf(best, next),
		next:    next,
		member:  func(spans []Span) bool { return cs.member(shiftSpans(spans, -delta)) },
		ways:    cs.ways,
		clashes: cs.clashes,
	}
	rg.start()
	return rg
}
