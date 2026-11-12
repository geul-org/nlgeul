//ff:func feature=token type=reader control=iteration dimension=2
//ff:what bundles — cut a unit at the agreed boundaries and gather, for each stretch, the different ways the surviving candidates read it, best first
package token

func (sc *scan) bundles(paths []*path) []Bundle {
	var out []Bundle
	from := 0
	for _, to := range agreed(paths) {
		b := Bundle{Text: string(sc.s[from:to])}
		seen := map[string]bool{}
		for _, p := range paths {
			br := stretch(p, from, to)
			if text := br.String(); !seen[text] {
				seen[text] = true
				b.Branches = append(b.Branches, br)
			}
		}
		out, from = append(out, b), to
	}
	return out
}
