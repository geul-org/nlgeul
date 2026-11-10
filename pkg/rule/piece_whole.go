//ff:func feature=rule type=rule control=sequence
//ff:what PieceWhole — 쓰임: counter. 수와 단위를 합친 글자열이 통째로 낱말 표의 낱말이거나 줄임 꼴이다(이건, 세대, 사회: 한 음절 한글 수사 뒤의 단위 풀이를 뒤로 보낸다). 기준 없음. 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceWhole(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Whole), nil
}
