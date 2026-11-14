//ff:func feature=token type=helper control=sequence
//ff:what stateBefore — the state before a candidate as the net names it: the scan's root for a candidate that has only the root (any object of it) before it, else the candidate before (a seed, put by a window in place of the root, is one)
package token

func (sc *scan) stateBefore(p *path) *path {
	if p.prev.prev == nil && !sc.registry().seeds[p.prev] {
		return sc.registry().root
	}
	return p.prev
}
