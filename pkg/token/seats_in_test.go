//ff:func feature=token type=helper control=sequence
//ff:what TestSeatsIn — each node becomes a seat of the region given, in order, reading as its first candidate; no node gives no seat
package token

import "testing"

func TestSeatsIn(t *testing.T) {
	seg, _ := sentenceRegions(t, "학생이되었다")[0].peek()
	seats := seatsIn(4, seg.Nodes)
	if len(seats) != 2 || seats[0].region != 4 || seats[1].region != 4 || seats[0].node.Span != (Span{0, 3}) || seats[1].node.Span != (Span{3, 6}) {
		t.Fatal(seats)
	}
	if seats[0].near != nearOf(seg.Nodes[0].Contest.Branches[0]) || seats[0].near.Last != "주격:이" || !seats[1].near.Pred {
		t.Fatal(seats[0].near, seats[1].near)
	}
	if len(seatsIn(0, nil)) != 0 {
		t.Fatal("no node")
	}
}
