//ff:func feature=rule type=rule control=sequence
//ff:what PieceLone — 쓰임: except. 조각 앞의 수가 홀로 선 한자어 한 음절 수사이고 그것을 가를 낱말 표가 없다(일, 이, 삼 …은 흔한 낱말의 첫 음절이다: 이건, 사회). 기준 없음. 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceLone(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Lone), nil
}
