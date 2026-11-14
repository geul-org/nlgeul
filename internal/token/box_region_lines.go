//ff:func feature=token type=helper control=selection
//ff:what boxRegionLines — the lines of one stretch of the stream from the ways taken for it (the first is the best): a stretch nothing is known of (no nodes) is its letters on one line, a stretch cut one way only is its boxes stacked, and more than one way is the boxes of the ways with the brackets (boxWaysLines)
package token

func boxRegionLines(segs []Segment, letters []rune, head int) []string {
	switch {
	case len(segs[0].Nodes) == 0:
		return []string{letterLine(segs[0].Region, letters)}
	case oneWay(segs[0]):
		return boxStackLines(segs[0], head)
	}
	return boxWaysLines(segs, head)
}
