//ff:func feature=token type=reader control=iteration dimension=1
//ff:what readGroup — judge every entry of one class that ends at j and merge the readings of the same token sequence at the same place in a chain into one piece (its names joined, its best verdict and count)
package token

import (
	"fmt"

	"github.com/geul-org/koparse/pkg/rule"
)

// The pieces come back by where their form starts, with one of the
// entries behind each. next is the place of the piece that follows.
func (sc *scan) readGroup(class, j, next int, more bool) (map[string]piece, map[string]*entry, map[string]int) {
	pieces, from, starts := map[string]piece{}, map[string]*entry{}, map[string]int{}
	for _, e := range sc.entriesAt(j) {
		start, p, ok := sc.place(e, j, more)
		if !ok || e.class != class {
			continue
		}
		p.Next = next
		v := sc.judge(e, start, j, p)
		if v <= 0 {
			continue
		}
		key := fmt.Sprintf("%s·%s·%d", rule.LeadText(e.lead), e.form, e.slot)
		pieces[key] = merged(pieces[key], piece{string(sc.s[start:j]), e.name, v, sc.ix.pref(e.count, e.form), tailOf(e), false})
		from[key], starts[key] = e, start
	}
	return pieces, from, starts
}
