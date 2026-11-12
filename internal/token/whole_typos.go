//ff:func feature=token type=reader control=iteration dimension=1
//ff:what wholeTypos — every candidate that ends at j with the letters from b as a word of no markers written wrong, once for each right form of the word table that the letters are a misspelt form of and that may stand with no marker after it, and once as words bound by a slash or a hyphen with a word written wrong read as its right form; grown from each candidate kept at b, as a known word of no markers is (the candidates that end in a modifying closing or a determiner are not grown by a word that only stands alone)
package token

func (sc *scan) wholeTypos(cands []*path, best [][]*path, b, j int) []*path {
	if len(best[b]) == 0 {
		return cands
	}
	reads := sc.typoReads(nil, string(sc.s[b:j]), tail{start: j})
	if sc.ix.boundWrong(string(sc.s[b:j])) {
		reads = append(reads, tail{start: j, boundTypo: true})
	}
	for _, t := range reads {
		for _, e := range sc.stepsOf(b, j, t) {
			cands = sc.growTypo(cands, best[b], t, e, b, j)
		}
	}
	return cands
}
