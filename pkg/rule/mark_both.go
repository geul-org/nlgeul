//ff:func feature=rule type=rule control=sequence
//ff:what MarkBoth — 쓰임: except. 기호의 양쪽이 모두 기준의 글자 갈래다(양쪽이 영숫자 열인 점은 문장을 끝내는 점이 아니고, 양쪽이 숫자인 빗금·붙임표는 이름을 잇는 것이 아니라 수의 기호다). 기준: Bits(sides). 값: 양쪽 다 맞으면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkBoth(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	m, want := MarkOf(ctx), specs[0].(*Bits).Mask
	return Applies(m.Left&want != 0 && m.Right&want != 0), nil
}
