//ff:func feature=rule type=rule control=sequence
//ff:what MarkRepeated — 쓰임: except. 점이 둘 이상 이어진다(마침표가 아니라 줄임표다). 기준 없음. 값: 이어지면 +1, 아니면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func MarkRepeated(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	return Applies(MarkOf(ctx).Run >= 2), nil
}
