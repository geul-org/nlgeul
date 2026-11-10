//ff:func feature=rule type=rule control=sequence
//ff:what ChainOrder — 쓰임: except. 조사 사슬에서 이 조사 뒤에 그 조사가 올 수 없다. 기준 없음. 값: 순서가 어긋나면(Follows가 거짓) +1, 맞거나 뒤가 없으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ChainOrder(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	return Applies(p.Next >= 0 && !Follows(p.Slot, p.Next)), nil
}
