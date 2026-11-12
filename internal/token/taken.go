//ff:func feature=token type=reader control=iteration dimension=1
//ff:what taken — the observed count of a key as it is read before a tail: that of the kind it stands as when a kind is given (the key's greatest when that kind has none); else the greatest among its kinds that take the tail, and the key's greatest when those kinds have none (a closed-set word the table lists under another kind), or when kinds are not looked at
package token

func (ix *index) taken(stem string, tail, kind int) int {
	w := ix.lex[stem]
	if ix.layers < layerKind {
		return w.count
	}
	best := 0
	for i, k := range kindOrder {
		if kind != 0 && k == kind {
			best = w.by[i]
		}
		if kind == 0 && w.kinds&k != 0 && kindTails[k]&tail != 0 {
			best = max(best, w.by[i])
		}
	}
	if best == 0 {
		return w.count
	}
	return best
}
