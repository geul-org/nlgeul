//ff:func feature=token type=helper control=iteration dimension=1
//ff:what headWalk — two streams of the same sentence walked side by side to the end, one with Next and Again and one with NextHead and AgainHead: every stretch and every way of it, what is wrong with each pair (headMatch), and whether both end at the same time
package token

import "fmt"

func headWalk(full, cut *Stream, head int) []string {
	var bad []string
	for {
		want, ok := full.Next()
		got, gotOK := cut.NextHead(head)
		if ok != gotOK {
			return append(bad, fmt.Sprintf("Next %v, NextHead %v", ok, gotOK))
		}
		if !ok {
			return bad
		}
		bad = append(bad, headMatch(want, got, head)...)
		bad = append(bad, headAgain(full, cut, want.Region, head)...)
	}
}
