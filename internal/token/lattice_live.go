//ff:func feature=token type=helper control=iteration dimension=1
//ff:what live — the states on a way that reads the whole unit, in the order they were made
package token

func (l *lattice) live() []*path {
	mark := l.reached()
	var out []*path
	for _, p := range l.all {
		if mark[p] {
			out = append(out, p)
		}
	}
	return out
}
