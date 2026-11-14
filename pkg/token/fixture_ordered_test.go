//ff:func feature=token type=helper control=iteration dimension=1
//ff:what orderedLines — the names (as a test writes them, each at the start of an indented line before a tab) that are missing from the printed lines or come before the one named just before them
package token

import "strings"

func orderedLines(lines string, names []string) []string {
	var bad []string
	last := -1
	for _, w := range names {
		at := strings.Index(lines, "  "+w+"\t")
		if at < 0 || at < last {
			bad = append(bad, w)
		}
		last = at
	}
	return bad
}
