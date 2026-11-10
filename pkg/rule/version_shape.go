//ff:func feature=rule type=rule control=sequence
//ff:what VersionShape — 쓰임: except. 버전의 꼴이 아니다: 마디가 모두 숫자(첫 마디는 v로 시작해도 됨)이고, v로 시작하거나 점이 둘 이상이어야 한다. 기준 없음. 값: 그 꼴이 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func VersionShape(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	m := MarkOf(ctx)
	return Applies(!m.Numeric || (!m.VHead && m.Dots < 2)), nil
}
