//ff:func feature=rule type=rule control=sequence
//ff:what PieceParticle — 쓰임: except. 조각 바로 뒤에 조사가 붙는다(꾸밈·맺음 뒤에는 조사가 못 붙는다: 사신이의 신은 시+ㄴ이 아니다). 기준 없음. 값: 조사가 붙으면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceParticle(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Right == RightParticle), nil
}
