//ff:func feature=parse type=rule control=sequence
//ff:what topicObject — 주장: 격조사 없이 보조사만 붙었다 → 보조사는 목적격 조사 자리에도 선다 → 목적어. 값: 그렇게 읽혔으면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func topicObject(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.size > 0 && c.kase == 0), c.last
}
