//ff:func feature=rule type=rule control=sequence
//ff:what PieceRight — 쓰임: except. 조각 바로 뒤에 선 것이 기준의 것이다(받침만 떼어 낸 꼴 뒤에 조사가 붙는다: 사신이의 신은 시+ㄴ이 아니다; 떼어 낸 받침이 뒤 글자와 표의 맺음을 이루지 않는다: 신고의 신; 단위 뒤에 글자가 그대로 이어진다: 사회적의 회는 단위가 아니다; 예외를 꺾는 예외로는 이름 꼴 맺음 바로 뒤에 조사가 붙는다: 문제기도). 기준: Bits(right) = 뒤에 선 것(Right…). 값: 기준의 것이면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceRight(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(PieceOf(ctx).Right&specs[0].(*Bits).Mask != 0), nil
}
