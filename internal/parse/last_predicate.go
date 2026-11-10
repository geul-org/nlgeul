//ff:func feature=parse type=rule control=sequence
//ff:what lastPredicate — 쓰임: rule. 마지막 어절에 용언 꼬리가 붙었다 → 서술어(입력 전제 1). 값: 그렇게 읽혔으면 +1
package parse

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func lastPredicate(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t := rule.TailOf(ctx)
	return rule.Applies(rule.WordOf(ctx).Last && t.Size > 0), t.Closing
}
