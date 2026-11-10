//ff:func feature=parse type=rule control=sequence
//ff:what conjunctiveWord — 주장: 접속 부사(그리고, 그러나, 그래서, 하지만)는 문장을 잇는 말로 다른 성분과 얽히지 않는다 → 독립어. 값: 어절이 기준 목록에 있으면 +1
package parse

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func conjunctiveWord(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	w := wordOf(ctx)
	return applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), w.text)), w.text
}
