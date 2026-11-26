//ff:func feature=rule type=rule control=sequence
//ff:what QuotedSentence — 쓰임: rule, except. 문장 안에서 판정하는 후보 바로 앞 어절이 닫는 기호(줄기 이름의 갈래 닫기)이고, 그 앞(문장을 맺는 부호 어절이 있으면 그것을 건너뛴 어절)의 마지막 조각이 맺음(이름이 맺음:으로 시작)이다: 닫는 기호가 맺음으로 끝난 인용 문장을 닫는다. 기준: List(문장을 맺는 부호의 이름 목록, 볼 사실은 ListPrevLast), 뒤집으려면 Not을 함께. 값: 그렇다면 +1, 아니면 -1(Not이면 반대). 증거: 맺음의 이름
package rule

import (
	"slices"
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func QuotedSentence(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	l := specs[0].(*List)
	before := AroundOf(ctx).Before
	if !strings.HasPrefix(nearAt(before, 0).Kind, "닫기:") {
		return Holds(false, specs), nil
	}
	i := 1
	if slices.Contains(l.Items, nearAt(before, i).Last) {
		i++
	}
	last := nearAt(before, i).Last
	return Holds(strings.HasPrefix(last, "맺음:"), specs), last
}
