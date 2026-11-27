//ff:func feature=data type=helper control=iteration dimension=1
//ff:what contentKind — what the characters of a value are, each kind once in sorted order joined by +
package main

import (
	"slices"
	"strings"
)

func contentKind(t *ucdTable, s string) string {
	seen := map[string]bool{}
	for _, r := range s {
		seen[runeKind(t, r)] = true
	}
	kinds := make([]string, 0, len(seen))
	for k := range seen {
		kinds = append(kinds, k)
	}
	slices.Sort(kinds)
	return strings.Join(kinds, "+")
}
