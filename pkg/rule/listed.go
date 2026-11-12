//ff:func feature=rule type=rule control=sequence
//ff:what Listed — 쓰임: rule, except, counter. 기준이 가리키는 사실(List…: 조각의 꼴, 앞 조각의 이름, 뒤 조각의 이름, 합쳐진 조각이 놓인 음절, 점 오른쪽 마디, 이름의 맨 오른쪽 마디, 어절, 꼬리의 맺음)이 기준 목록에 있다. 기준: List(목록과 볼 사실), 뒤집으려면 Not을 함께(목록에 없다). 값: 목록에 있으면 +1, 없으면 -1(Not이면 반대)
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func Listed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	fact := listedFact(ctx, l.Of)
	return Holds(slices.Contains(l.Items, fact), specs), fact
}
