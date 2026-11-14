//ff:func feature=token type=helper control=sequence
//ff:what TestSegmentHead — a head of 0 or less gives the segment as it is; a head of 2 keeps the first two candidates of a node that has more and every candidate of one that has fewer, with the place, rank, ways and lambda of the segment as they were; appending to what is given or changing it does not touch the segment it was cut from; a segment with no node is given as it is
package token

import (
	"math/big"
	"testing"
)

func TestSegmentHead(t *testing.T) {
	a, b, c := Branch{Verdict: 1, Pref: -1}, Branch{Verdict: 0.5, Pref: -2}, Branch{Verdict: 0.5, Pref: -3}
	seg := Segment{Region: Span{0, 4}, Rank: 1, Ways: big.NewInt(2), Lambda: -1, Nodes: []Node{
		{Span: Span{0, 2}, Contest: Contest{Text: "물을", Branches: []Branch{a, b, c}}},
		{Span: Span{2, 4}, Contest: Contest{Text: "마셔", Glued: true, Branches: []Branch{a}}},
	}}
	if got := segmentHead(seg, 0); &got.Nodes[0] != &seg.Nodes[0] || len(got.Nodes[0].Contest.Branches) != 3 {
		t.Fatal("0 gives all as they are")
	}
	if got := segmentHead(seg, -1); len(got.Nodes[0].Contest.Branches) != 3 {
		t.Fatal("less than 0 gives all")
	}
	got := segmentHead(seg, 2)
	if got.Region != seg.Region || got.Rank != 1 || got.Ways != seg.Ways || got.Lambda != -1 || len(got.Nodes) != 2 {
		t.Fatal(got)
	}
	first, second := got.Nodes[0], got.Nodes[1]
	if first.Span != (Span{0, 2}) || first.Contest.Text != "물을" || len(first.Contest.Branches) != 2 || first.Contest.Branches[1].Pref != -2 {
		t.Fatal(first)
	}
	if !second.Contest.Glued || len(second.Contest.Branches) != 1 || second.Contest.Branches[0].Pref != -1 {
		t.Fatal(second)
	}
	got.Nodes[0].Contest.Branches = append(got.Nodes[0].Contest.Branches, Branch{Pref: 9})
	got.Nodes[1].Contest.Text = "x"
	if len(seg.Nodes[0].Contest.Branches) != 3 || seg.Nodes[0].Contest.Branches[2].Pref != -3 || seg.Nodes[1].Contest.Text != "마셔" {
		t.Fatal("the segment it was cut from is touched", seg)
	}
	if none := segmentHead(Segment{Region: Span{0, 1}}, 1); none.Nodes != nil || none.Region != (Span{0, 1}) {
		t.Fatal(none)
	}
}
