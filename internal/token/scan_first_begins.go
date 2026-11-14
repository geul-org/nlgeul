//ff:func feature=token type=reader control=iteration dimension=1
//ff:what firstBegins — the places in the word of a bound stem between y and end (the word before the last, with a sign after it) from which the stem may begin that is farther than a new word's stem may be (before lo), and that begin a word that stands: the stem begins in the word, so only a part of it counts
package token

func (sc *scan) firstBegins(y, end, lo int) []int {
	var out []int
	for b := max(y, end-sc.ix.keyMax); b < min(end, lo); b++ {
		if !sc.inside[b] && sc.partFits(b, end) {
			out = append(out, b)
		}
	}
	return out
}
