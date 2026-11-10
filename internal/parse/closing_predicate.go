//ff:func feature=parse type=rule control=sequence
//ff:what closingPredicate — 주장: 종결·연결 맺음으로 끝난 용언 꼬리가 붙었다 → (문장 또는 이어진 절의) 서술어. 값: 그렇게 읽혔으면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func closingPredicate(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t := tailOf(ctx)
	return applies(t.kind&(kindFinal|kindLink) != 0), t.closing
}
