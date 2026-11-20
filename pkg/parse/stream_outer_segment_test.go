//ff:func feature=parse type=reader control=sequence
//ff:what TestStreamOuterSegment — a segment is given as it is when the stream has no table of places; with one its region and the span of each node are mapped on a copy of the nodes (the nodes of the segment kept are not touched), and a segment with no node stays with none
package parse

import "testing"

func TestStreamOuterSegment(t *testing.T) {
	seg := Segment{Region: Span{0, 5}, Rank: 2, Nodes: []Node{{Span: Span{0, 3}}, {Span: Span{3, 5}}}}
	plain := &Stream{}
	if got := plain.outerSegment(seg); got.Region != seg.Region || got.Nodes[1].Span != (Span{3, 5}) {
		t.Fatal(got)
	}
	st := &Stream{places: []int{0, 1, 2, 2, 3, 4}}
	got := st.outerSegment(seg)
	if got.Region != (Span{0, 4}) || got.Rank != 2 || got.Nodes[0].Span != (Span{0, 2}) || got.Nodes[1].Span != (Span{2, 4}) {
		t.Fatal(got)
	}
	if seg.Region != (Span{0, 5}) || seg.Nodes[0].Span != (Span{0, 3}) || seg.Nodes[1].Span != (Span{3, 5}) {
		t.Fatal("the kept segment was changed", seg)
	}
	if none := st.outerSegment(Segment{Region: Span{0, 3}}); none.Region != (Span{0, 2}) || none.Nodes != nil {
		t.Fatal(none)
	}
}
