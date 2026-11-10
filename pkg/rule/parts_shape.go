//ff:func feature=rule type=rule control=sequence
//ff:what PartsShape — 쓰임: except. 기호가 잇는 숫자 마디들이 기준 꼴이 아니다(25:70은 시각이 아니고 1,00은 자릿수가 아니다). 기준: 꼴(Shape). 값: 꼴이 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PartsShape(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	parts := MarkOf(ctx).Parts
	return Applies(len(parts) == 0 || !FitsShape(parts, specs[0].(*Shape))), nil
}
