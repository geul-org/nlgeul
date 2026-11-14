//ff:func feature=token type=helper control=iteration dimension=1
//ff:what jsonSegmentOf — a segment the stream sent as it is written in JSON; the letters of the stretch come from the letters of the sentence
package token

func jsonSegmentOf(seg Segment, letters []rune) JSONSegment {
	out := JSONSegment{
		Region: [2]int{seg.Region.Start, seg.Region.End},
		Text:   letterLine(seg.Region, letters),
		Rank:   seg.Rank,
		Lambda: jsonFloat(seg.Lambda),
		Ways:   jsonWays(seg.Ways),
		Words:  make([]JSONWord, 0, len(seg.Nodes)),
	}
	for _, n := range seg.Nodes {
		out.Words = append(out.Words, jsonWordOf(n))
	}
	return out
}
