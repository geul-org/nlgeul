//ff:func feature=token type=helper control=iteration dimension=1
//ff:what enclosedNumber — is r a circled or parenthesized number (①-⑳, ⑴-⒇, ㉑-㊿; enclosedRanges)
package token

func enclosedNumber(r rune) bool {
	for _, g := range enclosedRanges {
		if r >= g[0] && r <= g[1] {
			return true
		}
	}
	return false
}
