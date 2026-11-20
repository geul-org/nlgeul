//ff:func feature=parse type=reader control=iteration dimension=2
//ff:what contests — cut a unit at the agreed boundaries and gather, for each stretch, the different ways the surviving candidates read it, best first
package parse

func (sc *scan) contests(paths []*path) []Contest {
	var out []Contest
	from := 0
	for _, to := range agreed(paths) {
		b := Contest{Text: string(sc.s[from:to])}
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
