//ff:func feature=token type=helper control=iteration dimension=2
//ff:what stretchContest — a stretch the stream sent first (seg, its best way) set as the old structure: its letters, the unit and Glued of its first node, and as branches every way of it (seg, then Again until there is none) each as wayBranches gives it, a branch printed as one before left out
package token

func stretchContest(st *Stream, seg Segment, letters []rune) Contest {
	c := Contest{Text: string(letters[seg.Region.Start:seg.Region.End])}
	if len(seg.Nodes) > 0 {
		c.Unit, c.Glued = seg.Nodes[0].Contest.Unit, seg.Nodes[0].Contest.Glued
	}
	seen := map[string]bool{}
	for _, way := range streamWays(st, seg) {
		for _, b := range wayBranches(way) {
			if text := b.String(); !seen[text] {
				seen[text] = true
				c.Branches = append(c.Branches, b)
			}
		}
	}
	return c
}
