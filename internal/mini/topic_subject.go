//ff:func feature=mini type=rule control=sequence
//ff:what topicSubject — 주장: 격조사 없이 보조사만 붙었다(은/는, 도, 만 …) → 보조사는 주격 조사 자리에 선다 → 주어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func topicSubject(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	return applies(c.size > 0 && c.kase == 0), c.last
}
