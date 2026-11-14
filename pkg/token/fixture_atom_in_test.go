//ff:func feature=token type=helper control=iteration dimension=1
//ff:what atomIn — what atomStem gives for the stem from b to end as a test reads it: the names of its candidates joined with "," and the best verdict (an empty string and 0 for none)
package token

func atomIn(sc *scan, b, end int) (string, float64) {
	candidates := sc.atomStem(b, end)
	names := make([]kindCandidate, len(candidates))
	for i, r := range candidates {
		names[i] = kindCandidate{0, r.name, r.v}
	}
	return candidatesIn(names)
}
