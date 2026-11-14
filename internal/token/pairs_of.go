//ff:func feature=token type=helper control=iteration dimension=2
//ff:what pairsOf — the different pairs of words that stand side by side in some way of cutting, in the order they are met
package token

func pairsOf(cuts []cut) []edge {
	var out []edge
	seen := map[edge]bool{}
	for _, c := range cuts {
		for i := 1; i < len(c.spans); i++ {
			e := edge{from: c.spans[i-1], to: c.spans[i]}
			if !seen[e] {
				seen[e] = true
				out = append(out, e)
			}
		}
	}
	return out
}
