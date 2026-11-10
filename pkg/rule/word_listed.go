//ff:func feature=rule type=rule control=sequence
//ff:what WordListed — 쓰임: rule, except. 어절이 통째로 기준 목록의 낱말이다(관형사 → 관형어, 접속 부사 → 독립어; 예외로는 닫힌 집합의 낱말을 쪼개 읽지 않는다). 기준: 낱말 목록. 값: 목록에 있으면 +1, 아니면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func WordListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	w := WordOf(ctx)
	return Applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), w.Text)), w.Text
}
