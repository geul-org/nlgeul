//ff:func feature=rule type=rule control=sequence
//ff:what Longer — 쓰임: except. 경쟁 풀이가 이 풀이보다 어절 끝을 더 길게 읽는다(이 풀이는 긴 조각의 일부만 떼어 읽은 것이다). 기준: Versus(내 길이, 경쟁 길이). 값: 경쟁 쪽이 더 길면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func Longer(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	v := specs[0].(*Versus)
	return Applies(measure(ctx, v.Rival) > measure(ctx, v.Mine)), nil
}
