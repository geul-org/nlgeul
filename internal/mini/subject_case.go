//ff:func feature=mini type=rule control=sequence
//ff:what subjectCase — 주장: 주격 조사(이/가, 께서, 단체의 에서, 사람 수의 서)가 사슬의 가장 바깥 격조사다 → 주어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func subjectCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.kase&caseSubj != 0), c.last
}
