//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wordsOf — the words of a reading, first to last
package token

func wordsOf(p *path) []eojeol {
	var out []eojeol
	for ; p != nil && p.prev != nil; p = p.prev {
		out = append([]eojeol{p.last}, out...)
	}
	return out
}
