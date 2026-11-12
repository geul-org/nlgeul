//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordsOf — the words of a candidate, first to last
package token

func wordsOf(p *path) []step {
	var out []step
	for ; p != nil && p.prev != nil; p = p.prev {
		out = append([]step{p.last}, out...)
	}
	return out
}
