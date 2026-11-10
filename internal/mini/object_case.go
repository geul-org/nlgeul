//ff:func feature=mini type=rule control=sequence
//ff:what objectCase — 주장: 목적격 조사(을/를)가 사슬의 가장 바깥 격조사다 → 목적어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func objectCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.kase&caseObj != 0), c.last
}
