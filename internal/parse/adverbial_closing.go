//ff:func feature=parse type=rule control=sequence
//ff:what adverbialClosing — 주장: 맺음 게는 용언을 부사처럼 쓰이게 한다(morphemes.md 4절) → 부사어. 값: 읽힌 맺음이 기준 목록에 있으면 +1
package parse

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// adverbialClosing's criterion is the list of closings (adverbialClosings).
func adverbialClosing(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	t := tailOf(ctx)
	list := specs[0].(toulmin.StringLister).StringList()
	return applies(slices.Contains(list, t.closing)), t.closing
}
