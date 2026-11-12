//ff:func feature=rule type=rule control=sequence
//ff:what StemKind — 쓰임: except. 줄기가 표에서 기준 갈래 가운데 하나다(앞말이 표의 용언 줄기다, 표가 아는 낱말이다). 기준: 갈래 비트(Bits), 뒤집으려면 Not을 함께(앞말이 기준 갈래의 줄기가 아니다). 값: 기준 갈래가 하나라도 있으면 +1, 없으면 -1(Not이면 반대)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func StemKind(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Holds(PieceOf(ctx).Stem&specs[0].(*Bits).Mask != 0, specs), nil
}
