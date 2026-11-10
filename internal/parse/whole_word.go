//ff:func feature=parse type=rule control=sequence
//ff:what wholeWord — 예외: 닫힌 집합의 낱말(관형사, 접속 부사)은 통째로 한 낱말이다 → 조사나 용언 꼬리로 쪼개 읽지 않는다. 값: 어절이 기준 목록에 있으면 +1
package parse

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func wholeWord(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), wordOf(ctx).text)), nil
}
