//ff:func feature=rule type=rule control=sequence
//ff:what WordIs — 쓰임: except. 어절이 기준의 사실 가운데 하나를 갖는다(WordLast: 마지막 어절은 서술어다, WordBound: 의존명사는 용언 줄기가 아니다). 기준: Bits(word). 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func WordIs(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	w := WordOf(ctx)
	have := 0
	if w.Last {
		have |= WordLast
	}
	if w.Bound {
		have |= WordBound
	}
	return Applies(have&specs[0].(*Bits).Mask != 0), nil
}
