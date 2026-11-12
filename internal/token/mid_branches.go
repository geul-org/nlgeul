//ff:func feature=token type=helper control=iteration dimension=1
//ff:what midBranches — the branches a candidate becomes when each word that is a sign between letters (midName) is read by each candidate of that sign, the first sign's candidate changing slowest; a candidate with no such word is given back as it is
package token

func (sc *scan) midBranches(br Branch) []Branch {
	out := []Branch{br}
	for k, w := range br.Words {
		if m, ok := sc.midOf(w); ok {
			out = crossParts(out, k, m.candidates)
		}
	}
	return out
}
