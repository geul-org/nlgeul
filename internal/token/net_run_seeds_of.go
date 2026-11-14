//ff:func feature=token type=helper control=iteration dimension=1
//ff:what seedsOf — a seed for each root, in the order the roots end: a state of its own with the root's last word and nobody before it (so that it is a root as the unit's own is), noted in the registry; and which seed stands for which root
package token

func (r *netRun) seedsOf(order []*path, roots map[*path]bool) (map[*path]*path, []*path) {
	reg := r.sc.registry()
	seedOf, seeds := map[*path]*path{}, []*path{}
	reg.seeds = map[*path]bool{}
	for _, p := range order {
		if roots[p] {
			sd := &path{last: p.last, verdict: p.verdict, pref: p.pref}
			seedOf[p], reg.seeds[sd] = sd, true
			seeds = append(seeds, sd)
		}
	}
	return seedOf, seeds
}
