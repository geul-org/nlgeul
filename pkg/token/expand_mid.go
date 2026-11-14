//ff:func feature=token type=reader control=iteration dimension=2
//ff:what expandMid — the contests of a unit with each word that is a sign between letters (one part, named 기호) given the candidates of that sign, one branch for each candidate (every pairing when a branch holds several); the place of the sign is the end of its word less one; the order of the branches and what they hold are kept, and nothing is changed in the contests given
package token

func (sc *scan) expandMid(contests []Contest) []Contest {
	if len(sc.mid) == 0 {
		return contests
	}
	out := make([]Contest, len(contests))
	for i, b := range contests {
		out[i] = b
		out[i].Branches = nil
		for _, br := range b.Branches {
			out[i].Branches = append(out[i].Branches, sc.midBranches(br)...)
		}
	}
	return out
}
