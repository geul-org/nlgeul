//ff:func feature=token type=helper control=selection
//ff:what firstHead — what a syllable begins a tail with: ㄴ, ㅅ, and 네 (which begins with ㄴ too); nothing for another letter
package token

func firstHead(r rune) int {
	if r < '가' || r > '힣' {
		return 0
	}
	switch cho := int(r-'가') / 588; {
	case r == '네':
		return headN | headNe
	case cho == choN:
		return headN
	case cho == choS:
		return headS
	}
	return 0
}
