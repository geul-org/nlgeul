//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestSlotOrder — 시 before 었 before 겠 is in order, 었 may repeat, nothing after is in order; 겠 before 었 and 시 before 시 are not
package rule

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestSlotOrder(t *testing.T) {
	specs := toulmin.Specs{&Bits{Of: "repeat", Mask: 2}}
	cases := []struct {
		slot, next int
		want       float64
	}{{1, 2, -1}, {2, 3, -1}, {2, 2, -1}, {4, -1, -1}, {3, 2, 1}, {1, 1, 1}}
	for _, c := range cases {
		if v, _ := SlotOrder(pieceCtx(Piece{Slot: c.slot, Next: c.next}), specs); v != c.want {
			t.Errorf("%d %d: %v", c.slot, c.next, v)
		}
	}
}
