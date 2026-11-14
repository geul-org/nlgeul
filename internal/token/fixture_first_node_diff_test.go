//ff:func feature=token type=helper control=iteration dimension=1
//ff:what firstNodeDiff — where the first node that differs differs, as words ("" when no node does): a node differs by its span, unit, Glued, letters, or by its candidates: their words (as they are), verdict, preference and number of birth, in order
package token

import "fmt"

func firstNodeDiff(a, b []Node) string {
	for i := range a {
		if !sameNode(a[i], b[i]) {
			return fmt.Sprintf("%d %v %v | %v | %v", i, a[i].Span, b[i].Span, branchRows(a[i].Contest.Branches), branchRows(b[i].Contest.Branches))
		}
	}
	return ""
}
