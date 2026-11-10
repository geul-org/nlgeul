//ff:func feature=parse type=rule control=sequence
//ff:what beforeBound — 예외: 뒤 어절이 의존명사로 시작하고 이 어절이 꾸밈 맺음으로 읽힌다 → 의존명사를 꾸미는 말이다, 조사가 붙은 말이 아니다. 값: 그러면 +1
package parse

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func beforeBound(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := wordOf(ctx)
	adn := slices.ContainsFunc(w.tails, func(t tailFact) bool { return t.kind&kindAdn != 0 })
	return applies(w.nextBound && adn), nil
}
