//ff:func feature=rule type=rule control=sequence
//ff:what StemKind — 쓰임: except. 앞말이 낱말 표에서 기준 갈래의 줄기가 아니다. 기준: 갈래 비트(Bits). 값: 기준 갈래가 하나도 없으면 +1, 있으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func StemKind(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Stem&specs[0].(*Bits).Mask == 0), nil
}
