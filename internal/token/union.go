//ff:func feature=token type=reader control=iteration dimension=1
//ff:what union — two lists of names joined with "|", as one list without repeats
package token

import (
	"slices"
	"strings"
)

func union(a, b string) string {
	out := strings.Split(a, "|")
	if a == "" {
		out = nil
	}
	for _, n := range strings.Split(b, "|") {
		if !slices.Contains(out, n) {
			out = append(out, n)
		}
	}
	return strings.Join(out, "|")
}
