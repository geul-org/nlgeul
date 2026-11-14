//ff:func feature=token type=reader control=iteration dimension=1
//ff:what stretchOf — one portion of the net as a stretch of the unit the stream takes: its cells, one for each span where words stand (the candidates finished as readUnit finishes them: touched at the places where the unit begins, only for the first portion; a word printed the same once; best first by lambda and then by the number of birth), the pairs of spans, how many ways of cutting there are, and whether a line of spans is one
package token

func (sc *scan) stretchOf(u unit, pt portion, first bool) unitStretch {
	w := newWeave(sc.ix, pt)
	cs := cells{from: pt.from, edges: w.edges(), ways: w.ways(), member: w.member}
	for node, sp := range w.spans {
		bs, clashes := pooled(sc.touched(w.branches(node), u, first && sp.Start == 0))
		sortByLambda(bs)
		cs.nodes = append(cs.nodes, cell{span: sp, branches: bs})
		cs.clashes += clashes
	}
	cs.sets = w.sets
	return unitStretch{contest: Contest{Text: string(sc.s[pt.from:pt.to])}, cells: &cs}
}
