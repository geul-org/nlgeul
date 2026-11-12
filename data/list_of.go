//ff:func feature=data type=helper control=iteration dimension=1
//ff:what listOf — the names of a closed set in the order of their letters, joined with commas, for an error that tells what may be written
package data

import (
	"sort"
	"strings"
)

func listOf(names map[string]int) string {
	out := make([]string, 0, len(names))
	for name := range names {
		out = append(out, name)
	}
	sort.Strings(out)
	return strings.Join(out, ", ")
}
