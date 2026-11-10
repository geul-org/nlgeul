//ff:func feature=rule type=rule control=sequence
//ff:what MarkAfterWord — 쓰임: except. 기호의 왼쪽에 글자(영숫자 열이나 한글)가 붙어 있다(수의 맨 앞에 서는 기호가 아니다). 기준 없음. 값: 붙어 있으면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkAfterWord(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(MarkOf(ctx).Left != 0), nil
}
