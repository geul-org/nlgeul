//ff:func feature=rule type=rule control=sequence
//ff:what PieceAdn — 쓰임: except. 조각 바로 앞의 어절이 꾸밈 맺음으로 끝난다(꾸밈꼴은 체언을 꾸민다: 같은 단위에서 꾸밈 맺음 ㄴ·는·ㄹ·던 바로 뒤에 용언 줄기가 서지 못한다). 기준 없음. 값: 꾸밈 맺음이면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceAdn(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Adn), nil
}
