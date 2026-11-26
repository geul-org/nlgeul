//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what StemRun — 쓰임: except. 줄기가 영문·숫자 열이다(라틴 문자, 숫자, 밑줄뿐이다: 낱말 표의 줄보다 열 그대로 읽는다). 기준 없음. 값: 글자가 모두 열의 글자면 +1, 아니거나 비었으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func StemRun(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	form := PieceOf(ctx).Form
	for _, r := range form {
		if !RunRune(r) {
			return Applies(false), nil
		}
	}
	return Applies(form != ""), nil
}
