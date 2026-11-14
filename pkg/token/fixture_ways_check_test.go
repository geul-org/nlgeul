//ff:func feature=token type=helper control=iteration dimension=1
//ff:what waysCheck — every way a stream gives for a stretch tiles it (tiles) and holds the candidates it should: the names of those that do not
package token

func waysCheck(segs []Segment, letters []rune) []string {
	var bad []string
	for _, sg := range segs {
		if !tiles(sg, letters) {
			bad = append(bad, segmentLine(sg))
		}
	}
	return bad
}
