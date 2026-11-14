//ff:func feature=token type=helper control=iteration dimension=2
//ff:what spansOf — the different spans the ways of cutting put words at, by start and then end
package token

import "sort"

func spansOf(cuts []cut) []Span {
	var out []Span
	seen := map[Span]bool{}
	for _, c := range cuts {
		for _, sp := range c.spans {
			if !seen[sp] {
				seen[sp] = true
				out = append(out, sp)
			}
		}
	}
	sort.Slice(out, func(i, j int) bool { return spanBefore(out[i], out[j]) })
	return out
}
