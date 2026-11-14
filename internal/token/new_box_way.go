//ff:func feature=token type=helper control=iteration dimension=1
//ff:what newBoxWay — the boxes of the nodes of a segment (makeBox with head), not drawn yet
package token

func newBoxWay(seg Segment, head int) boxWay {
	w := boxWay{boxes: make([]box, 0, len(seg.Nodes))}
	for _, n := range seg.Nodes {
		w.boxes = append(w.boxes, makeBox(n, head))
	}
	return w
}
