//ff:func feature=token type=reader control=sequence
//ff:what fitsSurname — the proviso of article 48: the word before the boundary is a family name of the list with no part after
package token

import "slices"

func fitsSurname(g gap) bool {
	l := stemPart(g.left)
	return len(g.left.Parts) == 1 && l.Stem && slices.Contains(surnameWords, l.Text)
}
