//ff:func feature=rule type=rule control=sequence
//ff:what NextListed — 쓰임: except. 뒤에 오는 조각이 기준 목록에 있다(서술격 이는 는·은·을, 습니다, 동사 전용 맺음을 받지 않는다). 기준: 조각 이름 목록. 값: 목록에 있으면 +1, 없으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func NextListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	after := PieceOf(ctx).After
	return Applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), after)), after
}
