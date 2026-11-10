//ff:func feature=parse type=rule control=sequence
//ff:what chainSeenMore — 반박: 꼬리만큼 길게 읽히는 조사가 표에서 관측됐다 → 조사일 수 있다. 값: 두 관측 횟수 가운데 조사의 몫
package parse

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// A chain shorter than the tail says nothing about the syllables it
// does not read, so it does not rebut.
func chainSeenMore(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	t, c := tailOf(ctx), wordOf(ctx).chain
	if c.size < t.size {
		return -1, nil
	}
	return share(t.count, c.count), c.last
}
