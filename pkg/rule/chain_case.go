//ff:func feature=rule type=rule control=sequence
//ff:what ChainCase — 쓰임: rule. 사슬의 가장 바깥 격조사가 기준의 격 가운데 하나를 나타낸다(주격 → 주어, 목적격 → 목적어 …). 기준: Bits(case). 값: 그렇게 읽혔으면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ChainCase(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	c := WordOf(ctx).Chain
	return Applies(c.Kase&specs[0].(*Bits).Mask != 0), c.Last
}
