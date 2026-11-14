//ff:func feature=token type=helper control=iteration dimension=1
//ff:what headAgain — every way after the first of a stretch, asked of two streams of the same sentence, one with Again and one with AgainHead: what is wrong with each pair (headMatch), and whether both run out at the same time
package token

import "fmt"

func headAgain(full, cut *Stream, region Span, head int) []string {
	var bad []string
	for {
		want, ok := full.Again(region)
		got, gotOK := cut.AgainHead(region, head)
		if ok != gotOK {
			return append(bad, fmt.Sprintf("%v: Again %v, AgainHead %v", region, ok, gotOK))
		}
		if !ok {
			return bad
		}
		bad = append(bad, headMatch(want, got, head)...)
	}
}
