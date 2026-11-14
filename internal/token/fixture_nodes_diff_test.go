//ff:func feature=token type=helper control=iteration dimension=1
//ff:what nodesDiff — the first node of two stretches that differs (but for the numbers of birth), with its span and the candidates of both ("" when none does)
package token

import "fmt"

func nodesDiff(a, b *region) string {
	for k := range a.nodes {
		if !sameNodeLoose(a.nodes[k], b.nodes[k]) {
			return fmt.Sprint("node ", k, a.nodes[k].Span, " | ", branchRows(a.nodes[k].Contest.Branches), " | ", branchRows(b.nodes[k].Contest.Branches))
		}
	}
	return ""
}
