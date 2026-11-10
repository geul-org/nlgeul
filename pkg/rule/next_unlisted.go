//ff:func feature=rule type=rule control=sequence
//ff:what NextUnlisted — 쓰임: except. 뒤에 오는 조각이 기준 목록에 없다(이음 으는 으를 받는 꼬리 앞에만 선다: 시, 니, 니까, 나, 냐). 기준: 조각 이름 목록. 값: 목록에 없으면 +1, 있으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func NextUnlisted(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	after := PieceOf(ctx).After
	return Applies(!slices.Contains(specs[0].(toulmin.StringLister).StringList(), after)), after
}
