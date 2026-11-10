//ff:func feature=rule type=rule control=sequence
//ff:what MarkSides — 쓰임: except. 기호의 양쪽이 기준의 글자 갈래가 아니다(파일명·도메인의 점은 양쪽이 영숫자 열, 소수점은 양쪽이 숫자). 기준: Bits(sides). 값: 어느 한쪽이라도 아니면 +1, 둘 다 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkSides(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	m, want := MarkOf(ctx), specs[0].(*Bits).Mask
	return Applies(m.Left&want == 0 || m.Right&want == 0), nil
}
