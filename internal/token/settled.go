//ff:func feature=token type=reader control=iteration dimension=1
//ff:what settled — with a word table: when a reading made of known words only stands, the readings that lean on a new word are left out (the contest is tipped overwhelmingly); otherwise every reading stays
package token

func (ix *index) settled(paths []*path) []*path {
	if ix.layers == layerNone || len(paths) == 0 || paths[0].verdict <= ix.unknown {
		return paths
	}
	var out []*path
	for _, p := range paths {
		if p.verdict == paths[0].verdict {
			out = append(out, p)
		}
	}
	return out
}
