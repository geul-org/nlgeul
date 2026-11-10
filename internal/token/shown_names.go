//ff:func feature=token type=helper control=iteration dimension=1
//ff:what shownNames — the readings of a sign as an atom is named with them: without the sign, and a thousands separator as a number
package token

import "strings"

func shownNames(joined string) string {
	out := ""
	for _, n := range strings.Split(joined, "|") {
		label, _, _ := strings.Cut(n, ":")
		if as, ok := shownAs[label]; ok {
			label = as
		}
		out = union(out, label)
	}
	return out
}
