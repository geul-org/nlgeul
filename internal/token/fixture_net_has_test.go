//ff:func feature=token type=helper control=iteration dimension=3
//ff:what netHas — whether some candidate of some node of the stretches the stream reads (the net, not the old search that spreads every path) prints with the given text in it
package token

import "strings"

func netHas(r *Reader, s, text string) bool {
	for _, rg := range netRegions(r, s) {
		for _, n := range rg.nodes {
			for _, b := range n.Contest.Branches {
				if strings.Contains(b.String(), text) {
					return true
				}
			}
		}
	}
	return false
}
