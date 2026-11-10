//ff:func feature=rule type=rule control=sequence
//ff:what SeenFew — 쓰임: counter. 이 낱말은 관측이 적다(표에 있다는 것은 있을 수 있다는 것이지 그럴듯하다는 것이 아니다). 기준: Bits(enough) = 넉넉하다고 치는 횟수. 값: 적은 정도 1 - log(횟수)/log(넉넉한 횟수)를 [-1, 1]로; 횟수 0은 한 번과 같이 +1; 횟수를 모르면(음수) -1
package rule

import (
	"math"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// One observation, or a row of the table with none, is as few as can be
// (+1); the count named enough, or more, is not few at all (-1). A
// negative count is that of a word no table row gives one.
func SeenFew(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	count := PieceOf(ctx).Count
	if count < 0 {
		return -1, nil
	}
	few := 1 - math.Log10(float64(max(count, 1)))/math.Log10(float64(specs[0].(*Bits).Mask))
	return 2*math.Max(0, math.Min(1, few)) - 1, count
}
