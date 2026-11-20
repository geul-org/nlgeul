//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what reaches — whether a state stays when the net is kept from place q on: a state that begins at q when a root is before it, a state that begins after q when a state that stays is before it
package parse

func reaches(p *path, q int, roots, alive map[*path]bool) bool {
	for _, b := range p.preds {
		if (p.start == q && roots[b]) || (p.start > q && alive[b]) {
			return true
		}
	}
	return false
}
