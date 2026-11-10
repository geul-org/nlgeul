//ff:func feature=parse type=rule control=sequence
//ff:what beforeBound — 쓰임: except. 뒤 어절이 의존명사로 시작하고 이 어절이 꾸밈 맺음으로 읽힌다 → 의존명사를 꾸미는 말이다, 조사가 붙은 말이 아니다. 값: 그러면 +1
package parse

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func beforeBound(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := rule.WordOf(ctx)
	adn := slices.ContainsFunc(w.Tails, func(t rule.Tail) bool { return t.Kind&kindAdn != 0 })
	return rule.Applies(w.NextBound && adn), nil
}
