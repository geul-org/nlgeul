//ff:func feature=parse type=rule control=sequence
//ff:what boundStem — 예외: 의존명사는 용언 줄기가 아니다 → 의존명사 뒤의 은·는·을은 꾸밈 맺음이 아니라 조사다. 값: 어절이 의존명사로 시작하면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func boundStem(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return applies(wordOf(ctx).bound), nil
}
