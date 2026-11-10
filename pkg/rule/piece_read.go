//ff:func feature=rule type=rule control=sequence
//ff:what PieceRead — 쓰임: rule, counter. 색인에서 걸린 토큰 열이 이 자리에 있다(주장: 이 토큰 열은 이 표지다; 반박으로는 '열쇠 표에 없는 낱말이다'처럼 늘 서는 사정). 기준 없음. 값: 늘 +1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceRead(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return 1, PieceOf(ctx).Form
}
