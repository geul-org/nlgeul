//ff:func feature=parse type=rule control=sequence
//ff:what vocativeCase — 주장: 호격 조사(아/야, 이여, 이시여)가 붙었다 → 독립어. 값: 그렇게 읽혔으면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func vocativeCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.kase&caseVoc != 0), c.last
}
