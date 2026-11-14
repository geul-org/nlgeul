//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stemWith — some candidates of a stem as the candidates of the stem's letters (and the letters it is written with when they are wrong)
package token

func stemWith(stem, wrote string, candidates []kindCandidate) []stemCandidate {
	out := make([]stemCandidate, len(candidates))
	for i, r := range candidates {
		out[i] = stemCandidate{r, stem, wrote}
	}
	return out
}
