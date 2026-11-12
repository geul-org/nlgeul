//ff:func feature=token type=reader control=iteration dimension=1
//ff:what changedReads — a tail read once more for each stem of the word table that the stem letters are a changed form of, when the tail begins with what that change happens before and the stem takes the tail
package token

func (sc *scan) changedReads(out []tail, stem string, t tail) []tail {
	for _, a := range sc.ix.alts[stem] {
		if _, v := sc.ix.known(a.stem, tailClass(t.pieces), false); v > 0 && irrHeads[a.kind]&t.head != 0 && (t.irr == 0 || t.irr == a.kind) {
			read := t
			read.shown, read.irr = a.stem, 0
			out = append(out, read)
		}
	}
	return out
}
