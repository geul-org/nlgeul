//ff:func feature=mini type=rule control=sequence
//ff:what adverbialCase — 주장: 부사격 조사(에, 에서, 에게, 으로, 과/와, 보다, 처럼, 라고 …)가 사슬의 가장 바깥 격조사다 → 부사어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func adverbialCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.kase&caseAdv != 0), c.last
}
