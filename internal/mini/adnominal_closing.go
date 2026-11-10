//ff:func feature=mini type=rule control=sequence
//ff:what adnominalClosing — 주장: 꾸밈 맺음(ㄴ/은, 는, ㄹ/을, 던)으로 끝난 용언 꼬리가 붙었다 → 관형어. 값: 그렇게 읽혔으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func adnominalClosing(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t := tailOf(ctx)
	return applies(t.kind&kindAdn != 0), t.closing
}
