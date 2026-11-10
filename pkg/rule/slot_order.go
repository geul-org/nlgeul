//ff:func feature=rule type=rule control=sequence
//ff:what SlotOrder — 쓰임: except. 선어말 칸의 순서(시 < 었 < 겠 < 더 < 맺음)가 어긋난다; 었은 겹칠 수 있다. 기준: Bits(repeat) = 겹칠 수 있는 칸. 값: 어긋나면 +1, 맞거나 뒤가 없으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func SlotOrder(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	again := p.Slot == p.Next && p.Slot == specs[0].(*Bits).Mask
	return Applies(p.Next >= 0 && p.Slot >= p.Next && !again), nil
}
