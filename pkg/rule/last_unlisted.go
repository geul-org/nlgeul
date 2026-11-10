//ff:func feature=rule type=rule control=sequence
//ff:what LastUnlisted — 쓰임: except. 점으로 이어진 이름의 맨 오른쪽 마디가 기준 목록에 없다(도메인의 최상위 도메인). 기준: 마디 목록. 값: 없으면 +1, 있으면 -1
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func LastUnlisted(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	last := strings.ToLower(MarkOf(ctx).Last)
	return Applies(!slices.Contains(specs[0].(toulmin.StringLister).StringList(), last)), last
}
