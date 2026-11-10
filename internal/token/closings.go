//ff:func feature=token type=reader control=iteration dimension=1
//ff:what closings — every predicate tail that ends at j: a closing, alone or with the slots before it, and the same after the copula stem
package token

import "sort"

// more tells that 요 follows. A closing that sits in the stem's last
// syllable takes no slot before it. A closing that attaches only to
// the copula stem stands otherwise only after 아니.
func (sc *scan) closings(j int, more bool) []tail {
	if j < 1 {
		return nil
	}
	pieces, from, starts := sc.readGroup(classClosing, j, -1, more)
	keys := make([]string, 0, len(pieces))
	for k := range pieces {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var out []tail
	for _, k := range keys {
		out = append(out, sc.closingTails(pieces[k], from[k], starts[k])...)
	}
	out = append(out, sc.loneHonors(j, false)...)
	out = append(out, sc.negated(out)...)
	return append(out, sc.copulas(out)...)
}
