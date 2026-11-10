//ff:func feature=parse type=rule control=sequence
//ff:what complementShare — 쓰임: counter. 서술어가 되다·아니다이면 이/가가 붙은 말 가운데 하나는 주어가 아니라 보어다. 값: 이/가가 붙은 말 n개 가운데 이 말이 그 보어일 몫 1/n
package parse

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func complementShare(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	s := rule.WordOf(ctx).Sent
	if !needsComplement(s) || s.Nominative == 0 {
		return -1, nil
	}
	return 2/float64(s.Nominative) - 1, s.Stems
}
