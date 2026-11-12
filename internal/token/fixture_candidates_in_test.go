//ff:func feature=token type=helper control=iteration dimension=1
//ff:what candidatesIn — the names of some candidates of a stem joined with "," and the best verdict among them
package token

import "strings"

func candidatesIn(candidates []kindCandidate) (string, float64) {
	var names []string
	best := 0.0
	for _, r := range candidates {
		names, best = append(names, r.name), max(best, r.v)
	}
	return strings.Join(names, ","), best
}
