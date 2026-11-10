//ff:func feature=rule type=rule control=sequence
//ff:what ExtUnlisted — 쓰임: except. 점 바로 오른쪽 마디가 기준 목록에 없다(파일명의 확장자). 기준: 마디 목록. 값: 없으면 +1, 있으면 -1
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func ExtUnlisted(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	ext := strings.ToLower(MarkOf(ctx).Ext)
	return Applies(!slices.Contains(specs[0].(toulmin.StringLister).StringList(), ext)), ext
}
