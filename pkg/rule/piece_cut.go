//ff:func feature=rule type=rule control=sequence
//ff:what PieceCut — 쓰임: except. 떼어 낸 받침이 뒤 토큰과 표의 맺음을 이루지 않는다(단위 끝도 조사도 아닌 것이 뒤따른다: 신고의 신). 기준 없음. 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceCut(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Right == RightOther), nil
}
