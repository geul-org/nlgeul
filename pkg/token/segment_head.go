//ff:func feature=token type=helper control=iteration dimension=1
//ff:what segmentHead — a segment with the contest of each node cut to its first head candidates (in their lambda order, ties as they stand), a copy of its nodes so that the segment the stream keeps is not touched; all of them when head is 0 or less. Nothing else of the segment changes: its lambda is the sum of the first candidates, which stay
package token

// The cut slices have no room past head, so that appending to them does
// not write over the candidates the stream keeps.
func segmentHead(seg Segment, head int) Segment {
	if head <= 0 || len(seg.Nodes) == 0 {
		return seg
	}
	nodes := make([]Node, len(seg.Nodes))
	for i, n := range seg.Nodes {
		if len(n.Contest.Branches) > head {
			n.Contest.Branches = n.Contest.Branches[:head:head]
		}
		nodes[i] = n
	}
	seg.Nodes = nodes
	return seg
}
