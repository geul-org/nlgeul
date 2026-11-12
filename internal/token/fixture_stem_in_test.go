//ff:func feature=token type=helper control=iteration dimension=1
//ff:what stemIn — what stemCandidates gives for a tail as a test reads it: the stem of the first candidate, the names of all the candidates joined with ",", their best verdict, and the letters the stem is written with
package token

func stemIn(sc *scan, b int, t tail, sole bool) (string, string, float64, string) {
	candidates := sc.stemCandidates(b, t, sole)
	kinds := make([]kindCandidate, len(candidates))
	for i, r := range candidates {
		kinds[i] = r.kindCandidate
	}
	names, best := candidatesIn(kinds)
	return candidates[0].stem, names, best, candidates[0].wrote
}
