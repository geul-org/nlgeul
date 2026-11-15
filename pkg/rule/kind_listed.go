//ff:func feature=rule type=rule control=sequence
//ff:what KindListed — 쓰임: rule, except. 기준이 가리키는 사실(List…: 문장 안에서 판정하는 후보의 줄기 이름·마지막 조각의 이름, 바로 앞 어절의 마지막 조각의 이름 …)에서 쌍점 앞의 갈래(꾸밈:ㄹ의 꾸밈, 의존명사:수의 의존명사)가 기준 목록에 있다. 기준: List(갈래 목록과 볼 사실), 뒤집으려면 Not을 함께(목록에 없다). 값: 목록에 있으면 +1, 없으면 -1(Not이면 반대). 증거: 그 갈래
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func KindListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	kind, _, _ := strings.Cut(listedFact(ctx, l.Of), ":")
	return Holds(slices.Contains(l.Items, kind), specs), kind
}
