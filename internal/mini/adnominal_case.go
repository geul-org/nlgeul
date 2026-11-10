//ff:func feature=mini type=rule control=sequence
//ff:what adnominalCase — 주장: 관형격 조사(의)가 사슬의 가장 바깥 격조사다 → 관형어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func adnominalCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.kase&caseAdn != 0), c.last
}
