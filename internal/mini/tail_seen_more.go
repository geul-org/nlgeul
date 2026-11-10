//ff:func feature=mini type=rule control=iteration dimension=1
//ff:what tailSeenMore — 반박: 사슬만큼 길게 읽히는 용언 꼬리가 표에서 관측됐다 → 용언 꼬리일 수 있다. 값: 두 관측 횟수 가운데 꼬리의 몫(그런 꼬리 가운데 가장 큰 것)
package mini

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// A tail shorter than the chain does not rebut.
func tailSeenMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := wordOf(ctx)
	v := -1.0
	for _, t := range w.tails {
		if t.size >= w.chain.size {
			v = max(v, share(w.chain.count, t.count))
		}
	}
	return v, nil
}
