//ff:func feature=rule type=rule control=sequence
//ff:what AroundIs — 쓰임: rule, except. 문장 안에서 판정하는 후보와 그 둘레가 기준의 사실 가운데 하나를 갖는다(Around…: 후보의 어절이 바로 앞 어절에 빈칸 없이 붙어 있다, 바로 앞 어절이 체언으로 끝난다, 바로 뒤 어절이 꾸밈꼴 뒤에 설 수 있다(줄기가 체언·관형사이거나 꾸밈꼴·이름꼴로 끝난다), 바로 뒤 어절이 빈칸 없이 붙어 있다, 바로 앞 어절이 이름 없는 기호만의 열이다(가린 이름 ○○)). 기준: Bits(around), 뒤집으려면 Not을 함께. 값: 사실이 있으면 +1, 없으면 -1(Not이면 반대)
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func AroundIs(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Holds(aroundFacts(AroundOf(ctx))&specs[0].(*Bits).Mask != 0, specs), nil
}
