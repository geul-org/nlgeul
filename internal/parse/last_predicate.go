//ff:func feature=parse type=rule control=sequence
//ff:what lastPredicate — 주장: 마지막 어절에 용언 꼬리가 붙었다 → 서술어(입력 전제 1). 값: 그렇게 읽혔으면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func lastPredicate(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t := tailOf(ctx)
	return applies(wordOf(ctx).last && t.size > 0), t.closing
}
