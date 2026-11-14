//ff:func feature=token type=helper control=iteration dimension=1
//ff:what keepAlive — the states of the list that stay, in the order given (a list of its own)
package token

func keepAlive(ps []*path, alive map[*path]bool) []*path {
	out := make([]*path, 0, len(ps))
	for _, p := range ps {
		if alive[p] {
			out = append(out, p)
		}
	}
	return out
}
