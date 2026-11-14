//ff:func feature=token type=reader control=sequence
//ff:what fitsBound — article 42: the word after the boundary is a nominal whose stem is a dependent noun of the lists
package token

import "slices"

func fitsBound(g gap) bool {
	s := stemPart(g.right)
	return nominalStem(s) && slices.Contains(dependentWords, s.Text)
}
