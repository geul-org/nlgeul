//ff:func feature=rule type=rule control=sequence
//ff:what NextListed — 쓰임: except, counter. 뒤에 오는 조각이 기준 목록에 있다(서술격 이는 는·은·을, 습니다, 동사 전용 맺음을 받지 않는다; 빠진 이는 이름 꼴 맺음 기 앞에 서지 못한다; 예외를 꺾는 예외로는 한글 아닌 글자 뒤의 빠진 이가 야 앞에 선다). 기준: 조각 이름 목록. 값: 목록에 있으면 +1, 없으면 -1
package rule

import (
	"slices"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func NextListed(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	after := PieceOf(ctx).After
	return Applies(slices.Contains(specs[0].(toulmin.StringLister).StringList(), after)), after
}
