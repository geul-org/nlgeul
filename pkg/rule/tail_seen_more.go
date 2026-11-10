//ff:func feature=rule type=rule control=iteration dimension=1
//ff:what TailSeenMore — 쓰임: counter. 사슬만큼 길게 읽히는 용언 꼬리가 표에서 관측됐다(용언 꼬리일 수 있다). 기준 없음. 값: 두 관측 횟수 가운데 꼬리의 몫(그런 꼬리 가운데 가장 큰 것)을 [-1, 1]로; 없으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// A tail shorter than the chain does not rebut.
func TailSeenMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	w := WordOf(ctx)
	v := -1.0
	for _, t := range w.Tails {
		if t.Size >= w.Chain.Size {
			v = max(v, Share(w.Chain.Count, t.Count))
		}
	}
	return v, nil
}
