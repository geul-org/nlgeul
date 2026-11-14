//ff:func feature=token type=helper control=iteration dimension=1
//ff:what seedsAmong — the seeds of the roots that are among the states of a list (those the list kept at the place), in the order the roots end (order), so that a place keeps only the seeds of the states it kept before the cut
package token

func seedsAmong(order []*path, seedOf map[*path]*path, kept []*path) []*path {
	in := make(map[*path]bool, len(kept))
	for _, p := range kept {
		in[p] = true
	}
	var out []*path
	for _, p := range order {
		if sd, ok := seedOf[p]; ok && in[p] {
			out = append(out, sd)
		}
	}
	return out
}
