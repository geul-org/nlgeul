//ff:func feature=rule type=rule control=sequence
//ff:what PieceClash — 쓰임: counter. 어/아/여로 시작하는 조각이 앞 음절의 모음과 어긋난다(먹아, 잡어: 죽이지 않고 낮춘다). 기준 없음. 값: 어긋나면(Harmony가 거짓) +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceClash(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	form := []rune(p.Form)
	return Applies(len(form) > 0 && p.Prev != 0 && !Harmony(p.Prev, form[0])), nil
}
