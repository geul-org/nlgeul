//ff:func feature=mini type=rule control=sequence
//ff:what complementShare — 반박: 서술어가 되다·아니다이면 이/가가 붙은 말 가운데 하나는 주어가 아니라 보어다. 값: 이/가가 붙은 말 n개 가운데 이 말이 그 보어일 몫 1/n
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func complementShare(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	s := wordOf(ctx).sent
	if !needsComplement(s) || s.nominative == 0 {
		return -1, nil
	}
	return 2/float64(s.nominative) - 1, s.stems
}
