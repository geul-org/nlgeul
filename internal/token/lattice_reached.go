//ff:func feature=token type=helper control=iteration dimension=2
//ff:what reached — the states that are on a way that reads the whole unit: the states that end the unit and every state before them, found back by preds (the root is not one); the states that never come to the unit's end are not
package token

func (l *lattice) reached() map[*path]bool {
	mark := map[*path]bool{}
	stack := append([]*path{}, l.ends...)
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if mark[p] || p.prev == nil {
			continue
		}
		mark[p] = true
		stack = append(stack, p.preds...)
	}
	return mark
}
