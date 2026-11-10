//ff:func feature=rule type=rule control=sequence
//ff:what PieceNoStem — 쓰임: except. 표지 앞에 줄기가 없다(표지는 줄기 뒤에만 붙는다). 기준 없음. 값: 줄기가 없으면 +1, 있으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func PieceNoStem(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(!PieceOf(ctx).HasStem), nil
}
