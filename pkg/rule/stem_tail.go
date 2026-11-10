//ff:func feature=rule type=rule control=sequence
//ff:what StemTail — 쓰임: except. 줄기의 갈래가 뒤에 붙은 것과 맞지 않는다(체언 뒤에는 조사, 용언 줄기 뒤에는 선어말·맺음, 용언 줄기는 홀로 서지 못한다). 기준: Bits(tail) = 이 갈래가 받는 뒤(Tail…). 값: 맞지 않으면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func StemTail(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Tail&specs[0].(*Bits).Mask == 0), nil
}
