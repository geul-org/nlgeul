//ff:func feature=rule type=rule control=sequence
//ff:what PieceRunsOn — 쓰임: except. 조각 뒤에 글자가 그대로 이어진다(단위 뒤에는 조사 사슬이나 서술격 이다, 째가 오거나 거기서 단위가 끝난다: 사회적의 회는 단위가 아니다). 기준 없음. 값: 이어지면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceRunsOn(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Right == RightOther), nil
}
