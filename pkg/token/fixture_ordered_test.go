//ff:func feature=token type=helper control=iteration dimension=1
//ff:what orderedLines — the names (as a test writes them, each a whole line of streamText) that are missing from the lines or come before the one named just before them
package token

import "strings"

func orderedLines(lines string, names []string) []string {
	var bad []string
	last := -1
	for _, w := range names {
		at := strings.Index(lines, "\n"+w+"\n")
		if at < 0 || at < last {
			bad = append(bad, w)
		}
		last = at
	}
	return bad
}
