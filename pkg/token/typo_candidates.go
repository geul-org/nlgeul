//ff:func feature=token type=reader control=iteration dimension=1
//ff:what typoCandidates — a tail read once more for each right form that the stem letters are a misspelt form of in the word table, when the right form stands before the tail as a kind the table gives the misspelt form; the read shows the right form and keeps the letters
package token

func (sc *scan) typoCandidates(out []tail, stem string, t tail) []tail {
	for _, ty := range sc.ix.typos[stem] {
		if v := sc.ix.misspelt(ty.right, ty.kinds, tailClass(t.pieces), false); v > 0 {
			read := t
			read.shown, read.wrote, read.kinds = ty.right, stem, ty.kinds
			out = append(out, read)
		}
	}
	return out
}
