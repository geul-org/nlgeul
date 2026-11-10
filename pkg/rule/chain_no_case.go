//ff:func feature=rule type=rule control=sequence
//ff:what ChainNoCase — 쓰임: rule. 격조사 없이 보조사만 붙었다(보조사는 주격·목적격 조사 자리에 선다). 기준 없음(결론은 Claim이 가른다). 값: 그렇게 읽혔으면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ChainNoCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	c := WordOf(ctx).Chain
	return Applies(c.Size > 0 && c.Kase == 0), c.Last
}
