//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closedEnding — the closed-set words that end with a surface form: the exceptions of that form's graphs (그리고 for 고)
package token

import (
	"sort"
	"strings"
)

func closedEnding(words map[string]*entry, surface string) []string {
	var out []string
	for w := range words {
		if surface != "" && w != surface && strings.HasSuffix(w, surface) {
			out = append(out, w)
		}
	}
	sort.Strings(out)
	return out
}
