//ff:func feature=token type=reader control=sequence
//ff:what spacingNames — the names of the spacing rules that the boundary after word i is a place of, joined with "/" ("-" when there is none)
package token

import "strings"

func (sc *scan) spacingNames(words []step, i int) string {
	names := sc.ix.spacingOf(markGap(words, i))
	if len(names) == 0 {
		return "-"
	}
	return strings.Join(names, "/")
}
