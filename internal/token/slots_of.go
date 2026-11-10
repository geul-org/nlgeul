//ff:func feature=token type=reader control=iteration dimension=1
//ff:what slotsOf — the different kinds among some entries, each once, in order
package token

import "slices"

func slotsOf(entries []*entry) []int {
	var out []int
	for _, e := range entries {
		if !slices.Contains(out, e.slot) {
			out = append(out, e.slot)
		}
	}
	slices.Sort(out)
	return out
}
