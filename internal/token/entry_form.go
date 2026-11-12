//ff:func feature=token type=reader control=sequence
//ff:what entryForm — what an entry's count is kept under in the total: its class, its lead and its form, so that the candidates of one form (주격:이, 보격:이) are counted once
package token

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func entryForm(e *entry) string {
	return fmt.Sprintf("%d·%s·%s", e.class, rule.LeadText(e.lead), e.form)
}
