//ff:func feature=rule type=rule control=sequence
//ff:what ChainListed — 쓰임: rule. 격조사 없는 사슬 안에 기준 목록의 조사가 있다(범위의 까지·부터 → 부사어). 기준: 조사 목록. 값: 그러면 +1, 아니면 -1
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func ChainListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	c := WordOf(ctx).Chain
	list := specs[0].(toulmin.StringLister).StringList()
	found := slices.ContainsFunc(strings.Split(c.Parts, "+"), func(p string) bool { return slices.Contains(list, p) })
	return Applies(c.Kase == 0 && found), c.Last
}
