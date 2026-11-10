//ff:func feature=rule type=rule control=sequence
//ff:what ChainLinks — 쓰임: except. 조사 사슬에서 이 조사 뒤에 그 조사가 올 수 없다(앞 조사의 갈래가 뒤 조사의 갈래를 받지 않는다). 기준 없음. 값: 받지 않으면(Links가 거짓) +1, 받거나 뒤가 없으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// Slot and Next are the kinds (Link…) of this particle and of the one
// after it.
func ChainLinks(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	return Applies(p.Next >= 0 && !Links(p.Slot, p.Next)), nil
}
