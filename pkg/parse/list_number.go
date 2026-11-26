//ff:func feature=parse type=helper control=sequence
//ff:what listNumber — is a sign a round closing bracket (listBrackets) with a digit right before it: with none of its kind open it numbers an item of a list (1), 2))
package parse

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func listNumber(sign rune, m rule.Mark) bool {
	return m.Left&rule.SideDigit != 0 && strings.ContainsRune(listBrackets, sign)
}
