//ff:func feature=token type=reader control=sequence
//ff:what lambdaAt — the best lambda of the candidates at a span of the net: the candidates of the states there finished as stretchOf finishes them (touched, pooled, sortByLambda), the first one's lambda; kept for as long as the same number of states stand at the span
package token

func (r *netRun) lambdaAt(sp Span, states []*path) float64 {
	if got, ok := r.lam[sp]; ok && got.states == len(states) {
		return got.value
	}
	bs, _ := pooled(r.sc.touched(branchesOf(r.sc.ix, states), r.u, r.first && sp.Start == 0))
	sortByLambda(bs)
	v := lambdaOf(bs[0])
	r.lam[sp] = spanLambda{states: len(states), value: v}
	return v
}
