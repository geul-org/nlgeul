//ff:func feature=rule type=rule control=sequence
//ff:what DecimalShape — 쓰임: except. 소수의 꼴이 아니다: 수 사이의 점이 하나뿐이고 v로 시작하지 않아야 한다. 기준 없음. 값: 그 꼴이 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func DecimalShape(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	m := MarkOf(ctx)
	return Applies(m.VHead || m.Dots != 1), nil
}
