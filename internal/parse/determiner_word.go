//ff:func feature=parse type=rule control=sequence
//ff:what determinerWord — 주장: 지시·수 관형사(이, 그, 저, 이런, 어떤, 무슨, 한, 두)가 조사 없이 혼자 섰다 → 관형어. 값: 어절이 기준 목록에 있으면 +1
package parse

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func determinerWord(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	w := wordOf(ctx)
	return applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), w.text)), w.text
}
