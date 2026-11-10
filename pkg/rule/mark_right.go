//ff:func feature=rule type=rule control=sequence
//ff:what MarkRight — 쓰임: except. 기호의 오른쪽이 기준의 글자 갈래가 아니다(부호와 돈 기호는 수의 맨 앞에 선다). 기준: Bits(sides). 값: 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkRight(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(MarkOf(ctx).Right&specs[0].(*Bits).Mask == 0), nil
}
