//ff:func feature=rule type=rule control=sequence
//ff:what PieceNoNumber — 쓰임: except. 조각 바로 앞이 수가 아니다(단위 표의 낱말은 수 뒤에서만 단위다). 기준 없음. 값: 수가 아니면 +1, 수면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceNoNumber(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(!PieceOf(ctx).Num), nil
}
