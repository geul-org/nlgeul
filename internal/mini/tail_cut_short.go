//ff:func feature=mini type=rule control=sequence
//ff:what tailCutShort — 예외: 같은 끝에서 더 긴 꼬리가 읽힌다 → 이 풀이는 맺음을 중간에서 끊은 것이다. 값: 더 긴 꼬리가 있으면 +1
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func tailCutShort(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return longer(tailOf(ctx).size, wordOf(ctx).longest), nil
}
