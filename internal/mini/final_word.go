//ff:func feature=mini type=rule control=sequence
//ff:what finalWord — 예외: 마지막 어절은 서술어다(입력 전제 1) → 서술어가 아닌 성분으로 읽지 않는다. 값: 마지막 어절이면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func finalWord(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return applies(wordOf(ctx).last), nil
}
