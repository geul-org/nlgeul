//ff:func feature=token type=reader control=sequence
//ff:what prune — keep the net only from place q on, standing on the states xs (the last states of the best way of the stretch cut off: the states at its last span that a line of its spans leads to): the states that begin at q or later and can be reached from x are kept (survivors); when none can (blocked, which a cut at a place with a word after it from the states at the window's end does not give), every state that ends at q is a root (widened); the net is then put in order (reroot); no stretch is said to be unread here, that is told when the unit's end is reached (finish)
package token

func (r *netRun) prune(q int, xs []*path) {
	order := r.byEnd()
	roots := map[*path]bool{}
	for _, x := range xs {
		roots[x] = true
	}
	alive := survivors(order, q, roots)
	if len(alive) == 0 {
		r.stats.blocked++
		roots = endingAt(order, q)
		alive = survivors(order, q, roots)
		if len(alive) > 0 {
			r.stats.widened++
		}
	}
	r.reroot(q, roots, alive, order)
}
