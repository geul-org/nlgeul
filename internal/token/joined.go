//ff:func feature=token type=reader control=iteration dimension=1
//ff:what joined — each tail with one more piece on its right; the stem's restored syllable stays
package token

func joined(lefts []tail, pc piece) []tail {
	var out []tail
	for _, l := range lefts {
		ps := append(append([]piece{}, l.pieces...), pc)
		out = append(out, tail{start: l.start, pieces: ps, base: l.base, noYo: l.noYo})
	}
	return out
}
