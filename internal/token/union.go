//ff:func feature=token type=reader control=iteration dimension=1
//ff:what union — two lists of names as one list without repeats, in order
package token

import "slices"

func union(a, b []string) []string {
	out := slices.Clone(a)
	for _, n := range b {
		if !slices.Contains(out, n) {
			out = append(out, n)
		}
	}
	return out
}
