//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what SelfHas — 쓰임: rule, except. 문장 안에서 판정하는 후보의 표지(줄기 뒤 조각) 가운데 하나의 이름이 기준 목록에 있다. 기준: List(목록, 볼 사실은 ListSelfMarks), 뒤집으려면 Not을 함께. 값: 있으면 +1, 없으면 -1(Not이면 반대). 증거: 목록에 있던 이름
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func SelfHas(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	for _, m := range MarksOf(AroundOf(ctx).Self) {
		if slices.Contains(l.Items, m) {
			return Holds(true, specs), m
		}
	}
	return Holds(false, specs), nil
}
