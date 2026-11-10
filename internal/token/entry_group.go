//ff:func feature=token type=reader control=sequence
//ff:what entryGroup — the key readGroup files an entry's reading under: its lead, its form and its place or kind
package token

import (
	"fmt"

	"github.com/geul-org/koparse/pkg/rule"
)

func entryGroup(e *entry) string {
	return fmt.Sprintf("%s·%s·%d", rule.LeadText(e.lead), e.form, e.slot)
}
