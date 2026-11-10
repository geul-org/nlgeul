//ff:func feature=rule type=rule control=sequence
//ff:what PieceNoAdn — 쓰임: except. 조각 바로 앞의 어절이 꾸밈 맺음으로 끝나지 않는다(붙여 쓴 "것"의 줄임 게·건·걸은 꾸밈 맺음 ㄴ·는·ㄹ·던 바로 뒤에서만 읽는다: 한건, 하는게). 기준 없음. 값: 꾸밈 맺음이 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceNoAdn(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(!PieceOf(ctx).Adn), nil
}
