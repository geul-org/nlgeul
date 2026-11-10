//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TestChainOrder — 에+는, 만+으로, 는+요 are in order and a last particle has nothing after it; 이+는, 는+도, 는+에 are not
package rule

import "testing"

func TestChainOrder(t *testing.T) {
	cases := []struct {
		slot, next int
		want       float64
	}{{1, 4, -1}, {2, 1, -1}, {4, 5, -1}, {3, -1, -1}, {1, 1, -1}, {3, 4, 1}, {4, 4, 1}, {4, 1, 1}, {5, 4, 1}}
	for _, c := range cases {
		if v, _ := ChainOrder(pieceCtx(Piece{Slot: c.slot, Next: c.next}), nil); v != c.want {
			t.Errorf("%d %d: %v", c.slot, c.next, v)
		}
	}
}
