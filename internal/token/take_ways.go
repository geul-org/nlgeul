//ff:func feature=token type=helper control=iteration dimension=1
//ff:what takeWays — the ways of a stretch asked of the stream for the output, the best one (first) and then, one by one with Again, the ones after it up to head ways in all (0 for all); a stretch cut one way only is not asked for another, and no way beyond head is made
package token

func takeWays(st wayStream, first Segment, head int) []Segment {
	segs := []Segment{first}
	limit := waysLimit(first.Ways, head)
	for len(segs) < limit {
		seg, ok := st.Again(first.Region)
		if !ok {
			break
		}
		segs = append(segs, seg)
	}
	return segs
}
