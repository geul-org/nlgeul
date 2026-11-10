//ff:func feature=token type=reader control=iteration dimension=2
//ff:what bundles — cut a unit at the agreed boundaries and gather, for each stretch, the different ways the surviving readings read it, best first
package token

func (sc *scan) bundles(paths []*path) []bundle {
	var out []bundle
	from := 0
	for _, to := range agreed(paths) {
		b := bundle{text: string(sc.s[from:to])}
		seen := map[string]bool{}
		for _, p := range paths {
			br := stretch(p, from, to)
			if !seen[br.marked] {
				seen[br.marked] = true
				b.branches = append(b.branches, br)
			}
		}
		out, from = append(out, b), to
	}
	return out
}
