//ff:func feature=rule type=rule control=sequence
//ff:what MarkEnclosed — 쓰임: except. 점의 양쪽이 영숫자 열이다(문장을 끝내는 점이 아니라 이름을 잇는 점이다). 기준 없음. 값: 그러면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkEnclosed(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	m := MarkOf(ctx)
	return Applies(m.Left&SideAlnum != 0 && m.Right&SideAlnum != 0), nil
}
