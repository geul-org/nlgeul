//ff:func feature=rule type=rule control=sequence
//ff:what ChainSeenMore — 쓰임: counter. 판정 중인 꼬리만큼 길게 읽히는 조사가 표에서 관측됐다(조사일 수 있다). 기준 없음. 값: 두 관측 횟수 가운데 조사의 몫을 [-1, 1]로; 사슬이 더 짧거나 횟수를 모르면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// A chain shorter than the tail says nothing about the syllables it
// does not read, so it does not rebut.
func ChainSeenMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t, c := TailOf(ctx), WordOf(ctx).Chain
	if c.Size < t.Size {
		return -1, nil
	}
	return Share(t.Count, c.Count), c.Last
}
