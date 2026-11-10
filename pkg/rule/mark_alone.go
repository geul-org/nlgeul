//ff:func feature=rule type=rule control=sequence
//ff:what MarkAlone — 쓰임: except. 점이 하나뿐이다(줄임표는 점이 둘 이상 이어진다). 기준 없음. 값: 하나뿐이면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkAlone(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(MarkOf(ctx).Run < 2), nil
}
