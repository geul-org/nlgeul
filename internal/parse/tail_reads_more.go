//ff:func feature=parse type=rule control=sequence
//ff:what tailReadsMore — 예외: 어절 끝이 용언 꼬리로 더 길게 읽힌다 → 조사 풀이는 맺음의 끝 음절만 떼어 읽은 것이다. 값: 가장 긴 꼬리가 사슬보다 길면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func tailReadsMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := wordOf(ctx)
	return longer(w.chain.size, w.longest), nil
}
