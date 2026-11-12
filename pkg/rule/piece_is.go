//ff:func feature=rule type=rule control=sequence
//ff:what PieceIs — 쓰임: except, counter. 조각이 기준의 사실 가운데 하나를 갖는다(Piece…: 앞에 줄기가 있다, 바로 앞이 수다, 그 수가 홀로 선 한자어 한 음절 수사다, 수와 합쳐 통째로 낱말이다, 앞 어절이 꾸밈 맺음으로 끝난다, 그 어절 자신이 꾸밈꼴·이름꼴이다, 통짜 새 낱말이 표지로 읽힌다, 뒤에 이어지는 조각이 없다, 줄기가 끝 모음을 지킨다, 줄기의 낱말이 단위 전체다, 앞 어절이 관형사다). 기준: Bits(piece), 뒤집으려면 Not을 함께. 값: 사실이 있으면 +1, 없으면 -1(Not이면 반대)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceIs(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Holds(pieceFacts(PieceOf(ctx))&specs[0].(*Bits).Mask != 0, specs), nil
}
