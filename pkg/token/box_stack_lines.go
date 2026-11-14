//ff:func feature=token type=helper control=iteration dimension=1
//ff:what boxStackLines — the nodes of a segment of a stretch cut one way only (no contest of ways), each in a box of its own width, one over the other, with nothing at the right
package token

func boxStackLines(seg Segment, head int) []string {
	var out []string
	for _, n := range seg.Nodes {
		b := makeBox(n, head)
		out = append(out, b.lines(b.width)...)
	}
	return out
}
