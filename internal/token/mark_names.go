//ff:func feature=token type=helper control=iteration dimension=1
//ff:what markNames — the names of some candidates of a sign, in order
package token

func markNames(candidates []markCandidate) []string {
	out := make([]string, len(candidates))
	for i, r := range candidates {
		out[i] = r.name
	}
	return out
}
