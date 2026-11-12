//ff:func feature=token type=helper control=iteration dimension=1
//ff:what midBranches — the branches a reading becomes when each word that is a sign between letters (midName) is read by each reading of that sign, the first sign's reading changing slowest; a reading with no such word is given back as it is
package token

func (sc *scan) midBranches(br Branch) []Branch {
	out := []Branch{br}
	for k, w := range br.Words {
		if m, ok := sc.midOf(w); ok {
			out = crossParts(out, k, m.reads)
		}
	}
	return out
}
