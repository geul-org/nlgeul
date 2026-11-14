//ff:func feature=token type=helper control=iteration dimension=1
//ff:what shownNames — the candidates of a sign as an atom is named with them: without the sign, and a thousands separator as a number; a name that comes twice is one candidate, with the higher verdict
package token

import (
	"slices"
	"strings"
)

func shownNames(candidates []markCandidate) []markCandidate {
	var out []markCandidate
	for _, r := range candidates {
		label, _, _ := strings.Cut(r.name, ":")
		if as, ok := shownAs[label]; ok {
			label = as
		}
		i := slices.IndexFunc(out, func(o markCandidate) bool { return o.name == label })
		if i < 0 {
			out = append(out, markCandidate{label, r.v})
			continue
		}
		out[i].v = max(out[i].v, r.v)
	}
	return out
}
