//ff:func feature=parse type=rule control=sequence
//ff:what complementCase — 쓰임: rule. 서술어가 되다·아니다이고 이/가가 붙었다 → 보어. 값: 그렇게 읽혔으면 +1
package parse

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func complementCase(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := rule.WordOf(ctx)
	return rule.Applies(w.Chain.Kase&caseComp != 0 && needsComplement(w.Sent)), w.Chain.Last
}
