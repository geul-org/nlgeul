//ff:func feature=parse type=helper control=selection
//ff:what loneDepth — the depth of a sign that closes a pair with none of its kind open: a round bracket right after a digit numbers an item of a list and closes nothing (listBrackets), 0; any other closing bracket of loneBrackets closes the outermost (loneBracketDepth), and so does a quotation mark right after a mark that ends a sentence (MarkAfterStop: 다.”, the quotation began in a sentence before); any other quotation mark 0
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func loneDepth(sign rune, m rule.Mark) int {
	switch {
	case listNumber(sign, m):
		return 0
	case strings.ContainsRune(loneBrackets, sign) || m.Facts&rule.MarkAfterStop != 0:
		return loneBracketDepth
	}
	return 0
}
