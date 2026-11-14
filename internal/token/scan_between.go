//ff:func feature=token type=helper control=iteration dimension=1
//ff:what between — the states from place c to place q as a portion: the states that end at q and every state before them found back by preds (a seed or the root is none), in the order they were made
package token

func (sc *scan) between(c, q int) portion {
	reg := sc.registry()
	mark := map[*path]bool{}
	var stack []*path
	for _, p := range reg.all {
		if p.last.end == q {
			stack = append(stack, p)
		}
	}
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if mark[p] || p.prev == nil {
			continue
		}
		mark[p] = true
		stack = append(stack, p.preds...)
	}
	pt := portion{from: c, to: q}
	for _, p := range reg.all {
		if mark[p] {
			pt.states = append(pt.states, p)
		}
	}
	return pt
}
