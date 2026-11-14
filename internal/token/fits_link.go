//ff:func feature=token type=reader control=sequence
//ff:what fitsLink — article 45: the stem of the word after the boundary is a word that links or closes a list, or the word before it is such a word with no part after
package token

import "slices"

func fitsLink(g gap) bool {
	l := stemPart(g.left)
	return slices.Contains(linkWords, stemPart(g.right).Text) || (len(g.left.Parts) == 1 && l.Stem && slices.Contains(linkWords, l.Text))
}
