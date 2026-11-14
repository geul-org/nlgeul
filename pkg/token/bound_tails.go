//ff:func feature=token type=reader control=sequence
//ff:what boundTails — the candidates kept at b grown once more by the word that tail t makes when its stem is words bound by a slash or a hyphen with a word written wrong, read as the right form of that word (next to the candidate that takes the letters as they are, which variants gives)
package token

// Only a tail whose stem is the letters from b to its start as they
// stand is read so (not a fused or changed one, nor a stem a table
// names).
func (sc *scan) boundTails(cands []*path, kept []*path, t tail, b, j int) []*path {
	if t.base != 0 || t.shown != "" || t.as != "" || t.irr != 0 || !sc.ix.boundWrong(string(sc.s[b:t.start])) {
		return cands
	}
	t.boundTypo = true
	return sc.grown(cands, kept, t, b, j)
}
