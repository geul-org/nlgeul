//ff:func feature=rule type=rule control=sequence
//ff:what FormListed — 쓰임: except. 판정 중인 조각의 꼴이 기준 목록에 있다(의존 명사에 하다·싶다가 붙은 보조 용언의 줄기는 꾸밈 맺음 바로 뒤에 선다: 듯하, 만하). 기준: 꼴 목록. 값: 목록에 있으면 +1, 없으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func FormListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	form := PieceOf(ctx).Form
	return Applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), form)), form
}
