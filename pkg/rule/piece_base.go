//ff:func feature=rule type=rule control=sequence
//ff:what PieceBase — 쓰임: except. 줄기 끝 음절에 합쳐진 조각이 기준 목록에 없는 음절 밑에 있다(받침 하나뿐인 맺음은 하·되·시·키·이 밑에서만 읽는다). 기준: 음절 목록. 값: 목록에 없으면 +1, 있으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func PieceBase(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	return Applies(!slices.Contains(specs[0].(toulmin.StringLister).StringList(), PieceOf(ctx).Base)), nil
}
