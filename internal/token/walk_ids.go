//ff:func feature=token type=helper control=iteration dimension=1
//ff:what ids — the nodes of a walk, first to last
package token

func (w *walk) ids() []int {
	var out []int
	for p := w; p != nil; p = p.prev {
		out = append([]int{p.node}, out...)
	}
	return out
}
