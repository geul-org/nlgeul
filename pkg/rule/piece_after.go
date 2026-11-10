//ff:func feature=rule type=rule control=sequence
//ff:what PieceAfter — 쓰임: except. 앞 토큰의 끝이 이 표지가 붙는 조건과 맞지 않는다(은/는, 으로/로, 음/ㅁ). 기준: Bits(after) = 붙을 수 있는 끝. 값: 맞지 않으면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceAfter(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).PrevEnd&specs[0].(*Bits).Mask == 0), nil
}
