//ff:func feature=rule type=rule control=sequence
//ff:what FormUnlisted — 쓰임: except. 판정 중인 조각의 이름이 기준 목록에 없다(보조사가 뒤에 붙는 연결 어미는 지, 고, 어, 게다). 기준: 조각 이름 목록. 값: 목록에 없으면 +1, 있으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func FormUnlisted(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	form := PieceOf(ctx).Form
	return Applies(!slices.Contains(specs[0].(toulmin.StringLister).StringList(), form)), form
}
