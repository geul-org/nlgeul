//ff:func feature=rule type=rule control=sequence
//ff:what ShortKey — 쓰임: counter. 이 낱말은 짧아서 다른 말의 한 부분과 우연히 겹치기 쉽다. 기준: Bits(syllables) = 짧다고 치는 음절 수. 값: 그 길이 이하면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ShortKey(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	n := len([]rune(PieceOf(ctx).Form))
	return Applies(n > 0 && n <= specs[0].(*Bits).Mask), nil
}
