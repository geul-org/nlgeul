//ff:func feature=token type=reader control=iteration dimension=1
//ff:what reroot — put the net in order after a cut at q: each root becomes a seed (seedsOf), the states that stay are linked to the seeds and the states that stay before them (relink) and kept again in the registry in the order they were made (keep: the states, the edges, the numbers of birth, the words printed the same), and the places are put in order (trim)
package token

func (r *netRun) reroot(q int, roots, alive map[*path]bool, order []*path) {
	reg := r.sc.registry()
	seedOf, _ := r.seedsOf(order, roots)
	old := reg.all
	reg.all, reg.states, reg.preds, reg.born, reg.lone = nil, map[stateKey]*path{}, map[predKey]bool{}, map[wordKey]int{}, map[wordKey]bool{}
	for _, p := range old {
		if alive[p] {
			relink(p, q, roots, alive, seedOf)
			reg.keep(p)
		}
	}
	r.trim(q, seedOf, order, alive)
}
