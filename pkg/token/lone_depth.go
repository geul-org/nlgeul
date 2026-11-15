//ff:func feature=token type=helper control=sequence
//ff:what loneDepth — the depth of a sign that closes a pair with none of its kind open: a closing bracket of loneBrackets closes the outermost (loneBracketDepth), a quotation mark 0
package token

import "strings"

func loneDepth(sign rune) int {
	if strings.ContainsRune(loneBrackets, sign) {
		return loneBracketDepth
	}
	return 0
}
