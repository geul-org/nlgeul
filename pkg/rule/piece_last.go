//ff:func feature=rule type=rule control=sequence
//ff:what PieceLast — 쓰임: except. 이 표지 뒤에 이어지는 조각이 없다(뒤에 조사가 이어질 때만 읽는 표지: 받침으로 끝난 이름 뒤의 이). 기준 없음. 값: 뒤가 없으면 +1, 있으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceLast(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Next < 0), nil
}
