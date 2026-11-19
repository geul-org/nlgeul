//ff:func feature=token type=helper control=sequence
//ff:what loneDepth — the depth of a sign that closes a pair with none of its kind open: a closing bracket of loneBrackets closes the outermost (loneBracketDepth), and so does a quotation mark right after a mark that ends a sentence (MarkAfterStop: 다.”, the quotation began in a sentence before); any other quotation mark 0
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func loneDepth(sign rune, facts int) int {
	if strings.ContainsRune(loneBrackets, sign) || facts&rule.MarkAfterStop != 0 {
		return loneBracketDepth
	}
	return 0
}
