//ff:func feature=token type=reader control=sequence
//ff:what entryGroup — the key readGroup files an entry's reading under: its lead, its form and its place or kind, and for a closing written wrong the letters it is written with (it is never merged with the right closing of the same letters)
package token

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func entryGroup(e *entry) string {
	key := fmt.Sprintf("%s·%s·%d", rule.LeadText(e.lead), e.form, e.slot)
	if e.wrote != "" {
		key += "·오표기:" + e.wrote
	}
	return key
}
