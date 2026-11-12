//ff:func feature=token type=reader control=iteration dimension=2
//ff:what expandMid — the bundles of a unit with each word that is a sign between letters (one part, named 기호) given the candidates of that sign, one branch for each candidate (every pairing when a branch holds several); the place of the sign is the end of its word less one; the order of the branches and what they hold are kept, and nothing is changed in the bundles given
package token

func (sc *scan) expandMid(bundles []Bundle) []Bundle {
	if len(sc.mid) == 0 {
		return bundles
	}
	out := make([]Bundle, len(bundles))
	for i, b := range bundles {
		out[i] = b
		out[i].Branches = nil
		for _, br := range b.Branches {
			out[i].Branches = append(out[i].Branches, sc.midBranches(br)...)
		}
	}
	return out
}
