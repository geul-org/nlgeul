//ff:func feature=rule type=rule control=sequence
//ff:what MarkIs — 쓰임: except. 기호가 기준의 사실 가운데 하나를 갖는다(Mark…: 점이 둘 이상 이어진다, 왼쪽에 글자가 붙어 있다, 이름의 마디가 모두 숫자다). 기준: Bits(mark), 뒤집으려면 Not을 함께(점이 하나뿐이다, 마디가 모두 숫자가 아니다). 값: 사실이 있으면 +1, 없으면 -1(Not이면 반대)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkIs(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Holds(markFacts(MarkOf(ctx))&specs[0].(*Bits).Mask != 0, specs), nil
}
