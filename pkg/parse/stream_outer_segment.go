//ff:func feature=parse type=helper control=iteration dimension=1
//ff:what outerSegment — a segment as it is sent up: its region and the span of each node in the places of the layer above (outer); the nodes are copied so that the segment the region keeps is not touched, and the segment is the same when no unit holds a space inside a span of code
package parse

func (s *Stream) outerSegment(seg Segment) Segment {
	if s.places == nil {
		return seg
	}
	seg.Region = s.outer(seg.Region)
	if seg.Nodes == nil {
		return seg
	}
	nodes := make([]Node, len(seg.Nodes))
	for i, n := range seg.Nodes {
		n.Span = s.outer(n.Span)
		nodes[i] = n
	}
	seg.Nodes = nodes
	return seg
}
