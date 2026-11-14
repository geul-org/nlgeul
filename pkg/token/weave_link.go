//ff:func feature=token type=helper control=iteration dimension=2
//ff:what link — for each state the states before it that are in the portion, and for each span the spans that follow it, each once, the nearest first
package token

import "sort"

func (w *weave) link(idOf map[*path]int) {
	w.preds = make([][]int, len(w.states))
	w.succ = make([][]int, len(w.spans))
	seen := map[[2]int]bool{}
	for id, p := range w.states {
		for _, r := range p.preds {
			rid, ok := idOf[r]
			if !ok {
				continue
			}
			w.preds[id] = append(w.preds[id], rid)
			if e := [2]int{w.node[rid], w.node[id]}; !seen[e] {
				seen[e] = true
				w.succ[e[0]] = append(w.succ[e[0]], e[1])
			}
		}
	}
	for _, ids := range w.succ {
		sort.Ints(ids)
	}
}
