//ff:func feature=rule type=rule control=sequence
//ff:what ClosingKind — 쓰임: rule. 판정 중인 꼬리의 맺음이 기준의 갈래 가운데 하나다(종결·연결 → 서술어, 꾸밈 → 관형어). 기준: Bits(kind). 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ClosingKind(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	t := TailOf(ctx)
	return Applies(t.Kind&specs[0].(*Bits).Mask != 0), t.Closing
}
