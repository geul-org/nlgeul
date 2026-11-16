//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what PrevHas — 쓰임: rule, except. 문장 안에서 판정하는 후보 바로 앞 어절의 표지(줄기 뒤 조각) 가운데 하나의 이름이 기준 목록에 있다. 기준: List(목록, 볼 사실은 ListPrevLast — 앞 어절), 뒤집으려면 Not을 함께. 값: 있으면 +1, 없으면 -1(Not이면 반대). 앞 어절이 없으면 없다. 증거: 목록에 있던 이름
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func PrevHas(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	for _, m := range MarksOf(nearAt(AroundOf(ctx).Before, 0)) {
		if slices.Contains(l.Items, m) {
			return Holds(true, specs), m
		}
	}
	return Holds(false, specs), nil
}
