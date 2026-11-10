//ff:func feature=rule type=rule control=sequence
//ff:what ClosingListed — 쓰임: rule. 판정 중인 꼬리의 맺음이 기준 목록에 있다(게 → 부사어). 기준: 맺음 목록. 값: 그러면 +1, 아니면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func ClosingListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	t := TailOf(ctx)
	return Applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), t.Closing)), t.Closing
}
