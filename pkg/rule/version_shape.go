//ff:func feature=rule type=rule control=sequence
//ff:what VersionShape — 쓰임: except. 버전의 꼴이 아니다: 마디가 모두 숫자여야 한다(첫 마디는 v로 시작해도 됨). 점이 하나뿐인 수(1.2)는 소수점과 버전의 경합이다. 기준 없음. 값: 그 꼴이 아니면 +1, 맞으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func VersionShape(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	m := MarkOf(ctx)
	return Applies(!m.Numeric), nil
}
