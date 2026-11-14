//ff:func feature=token type=helper control=iteration dimension=1
//ff:what jsonCandidatesAll — whether every word of a stretch set down as data has as many candidates as the node the stream sent (none cut)
package token

func jsonCandidatesAll(js JSONSegment, seg Segment) bool {
	for i, n := range seg.Nodes {
		if len(js.Words[i].Candidates) != len(n.Contest.Branches) {
			return false
		}
	}
	return true
}
