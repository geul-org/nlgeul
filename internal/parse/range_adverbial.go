//ff:func feature=parse type=rule control=sequence
//ff:what rangeAdverbial — 주장: 격조사 없이 범위의 보조사(까지, 부터)가 붙었다 → 범위를 나타내는 부사어로도 쓰인다 → 부사어. 값: 그렇게 읽혔으면 +1
package parse

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// rangeAdverbial's criterion is the list of particles (rangeParticles).
func rangeAdverbial(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	c := wordOf(ctx).chain
	list := specs[0].(toulmin.StringLister).StringList()
	found := slices.ContainsFunc(strings.Split(c.parts, "+"), func(p string) bool { return slices.Contains(list, p) })
	return applies(c.kase == 0 && found), c.last
}
