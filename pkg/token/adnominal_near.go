//ff:func feature=token type=helper control=sequence
//ff:what adnominalNear — is a candidate, as a rule of a sentence reads it, an adnominal (관형어): its last part is a modifying closing, or it is a determiner's stem (관형사:이, 열쇠:관형사) standing alone with no marker
package token

import (
	"strings"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func adnominalNear(n rule.Near) bool {
	if strings.HasPrefix(n.Last, adnPrefix) {
		return true
	}
	return n.Marks == "" && (n.Kind == keyKind+":"+determinerLabel || strings.HasPrefix(n.Kind, determinerLabel+":"))
}
