//ff:func feature=parse type=rule control=sequence
//ff:what chainReadsMore — 예외: 어절 끝이 조사 사슬로 더 길게 읽힌다 → 꼬리 풀이는 사슬의 뒤쪽만 떼어 읽은 것이다. 값: 사슬이 더 길면 +1
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func chainReadsMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return longer(tailOf(ctx).size, wordOf(ctx).chain.size), nil
}
