//ff:func feature=token type=helper control=sequence
//ff:what TestRegionSegment — nodes made a segment: the region, the rank after those made, Ways, the nodes in order and the sum of their best lambdas
package token

import "testing"

func TestRegionSegment(t *testing.T) {
	rg := skewRegion()
	seg := rg.segment([]int{0, 2, 3, 5})
	if seg.Region != (Span{0, 5}) || seg.Rank != 1 || seg.Ways.Int64() != 2 || seg.Lambda != -3.5 || segmentLine(seg) != "가[0-1] 나[1-2] 다[2-3] 라마[3-5]" {
		t.Fatal(seg, segmentLine(seg))
	}
	rg.out = make([]Segment, 4)
	if rg.segment([]int{1}).Rank != 5 {
		t.Fatal("the rank follows those made")
	}
}
