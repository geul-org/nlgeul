//ff:func feature=rule type=rule control=sequence
//ff:what PieceMarked — 쓰임: except. 단위 안의 경계 뒤에 선 통짜 새 낱말이 통째로 색인의 표지(조사 사슬, 선어말과 맺음)로 읽힌다(조사와 어미는 앞말에 붙는다: 선생님 뒤의 께서는 새 낱말이 아니다). 기준 없음. 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceMarked(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Marked), nil
}
