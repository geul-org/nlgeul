//ff:func feature=token type=helper control=iteration dimension=1
//ff:what hasBranch — whether one of the branches prints as the given line
package token

func hasBranch(bs []Branch, line string) bool {
	for _, b := range bs {
		if b.String() == line {
			return true
		}
	}
	return false
}
