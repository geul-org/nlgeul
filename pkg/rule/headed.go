//ff:func feature=rule type=rule control=sequence
//ff:what Headed — 쓰임: rule, except. 문장 안에서 판정하는 후보 앞에 어절이 없거나(문장 첫머리), 바로 앞 어절의 줄기 이름의 갈래(열기:1의 열기 — 여닫는 기호는 줄기 자리에 선다)가 기준 목록에 있다. 줄기 이름이 없는 새 낱말은 첫머리가 아니다. 기준: List(갈래 목록; 볼 사실은 ListPrevKind), 뒤집으려면 Not을 함께. 값: 그렇다면 +1, 아니면 -1(Not이면 반대). 증거: 바로 앞 어절의 줄기 이름(첫머리면 "")
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func Headed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	before := AroundOf(ctx).Before
	if len(before) == 0 {
		return Holds(true, specs), ""
	}
	kind, _, _ := strings.Cut(before[0].Kind, ":")
	return Holds(kind != "" && slices.Contains(l.Items, kind), specs), before[0].Kind
}
