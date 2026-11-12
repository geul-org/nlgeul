//ff:func feature=rule type=rule control=sequence
//ff:what StemCoda — 쓰임: except. 뒤에 표지가 붙은 줄기의 마지막 음절이 기준 받침으로 끝난다(ㅆ 받침으로 끝나는 줄기는 있 계열뿐이고 나머지 ㅆ은 선어말 았·었·겠이다: 표에 없는 줄기는 ㅆ으로 끝나지 못한다). 기준: Bits(jong) = 받침 번호(Jong…). 값: 표지가 붙었고 그 받침으로 끝나면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func StemCoda(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	r := []rune(p.Form)
	return Applies(len(r) > 0 && p.Tail&^TailBare != 0 && JongOf(r[len(r)-1]) == specs[0].(*Bits).Mask), p.Form
}
