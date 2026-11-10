//ff:func feature=token type=reader control=iteration dimension=1
//ff:what taken — the observed count of a key as it is read before a tail: the greatest among its kinds that take the tail; the key's greatest when those kinds have none (a closed-set word the table lists under another kind), or when kinds are not looked at
package token

func (ix *index) taken(stem string, tail int) int {
	w := ix.lex[stem]
	if ix.layers < layerKind {
		return w.count
	}
	best := 0
	for i, k := range kindOrder {
		if w.kinds&k != 0 && kindTails[k]&tail != 0 {
			best = max(best, w.by[i])
		}
	}
	if best == 0 {
		return w.count
	}
	return best
}
