//ff:func feature=token type=helper control=iteration dimension=1
//ff:what pathOf — the letters of each node of a way, in order; a stretch with no node is its letters as they are
package token

func pathOf(seg Segment, letters []rune) []string {
	if len(seg.Nodes) == 0 {
		return []string{letterLine(seg.Region, letters)}
	}
	out := make([]string, 0, len(seg.Nodes))
	for _, n := range seg.Nodes {
		out = append(out, n.Contest.Text)
	}
	return out
}
