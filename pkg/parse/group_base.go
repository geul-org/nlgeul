//ff:func feature=parse type=reader control=sequence
//ff:what groupBase — the key an entry's place or kind is filed under: its lead, its form and its place or kind, and for a closing written wrong the letters it is written with (it is never merged with the right closing of the same letters), and for a particle read against the norm (misfit) misfitKey (it is never merged with a particle of the same name and form read only after a closing), and for a closing read only after a word that passes a graph (gate: a closing of the table of gated closings) the name of that graph (it is never merged with a closing of the same letters and name read with no such graph: 야 after the copula stem and the old spelling 야 of 여 after 하)
package parse

import (
	"fmt"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func groupBase(e *entry) string {
	key := fmt.Sprintf("%s·%s·%d", rule.LeadText(e.lead), e.form, e.slot)
	if e.wrote != "" {
		key += "·오표기:" + e.wrote
	}
	if e.misfit {
		key += "·" + misfitKey
	}
	if e.class == classClosing && e.gate != nil {
		key += "·" + e.gate.name
	}
	return key
}
