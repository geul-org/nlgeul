//ff:func feature=token type=helper control=iteration dimension=1
//ff:what shownNames — the candidates of a sign as an atom is named with them: without the sign, and a thousands separator as a number; a name that comes twice is one candidate, with the higher verdict
package token

import (
	"slices"
	"strings"
)

func shownNames(reads []markRead) []markRead {
	var out []markRead
	for _, r := range reads {
		label, _, _ := strings.Cut(r.name, ":")
		if as, ok := shownAs[label]; ok {
			label = as
		}
		i := slices.IndexFunc(out, func(o markRead) bool { return o.name == label })
		if i < 0 {
			out = append(out, markRead{label, r.v})
			continue
		}
		out[i].v = max(out[i].v, r.v)
	}
	return out
}
