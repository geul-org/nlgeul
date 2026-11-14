//ff:func feature=token type=helper control=iteration dimension=1
//ff:what more — the next way of the region in the order of lambda, greatest first: the best walk is taken out of the frontier and taken further until one reaches the end of the region and is one of the ways of cutting there are (best first search: f is the most a walk can still reach, so the ways come out in the order of their lambda; of equal lambda in the order the walks were put in); false when no way is left
package token

func (rg *region) more() (Segment, bool) {
	for len(rg.front.items) > 0 {
		w := rg.front.pop()
		if len(rg.next[w.node]) > 0 {
			rg.extend(w)
			continue
		}
		if seg, ok := rg.take(w); ok {
			return seg, true
		}
	}
	return Segment{}, false
}
