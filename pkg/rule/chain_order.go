//ff:func feature=rule type=rule control=sequence
//ff:what ChainOrder — 쓰임: except. 조사 사슬에서 이 조사 뒤에 그 조사가 올 수 없다(자리는 왼쪽에서 오른쪽으로 오른다; 격조사 뒤에는 요만, 범위 조사 뒤에는 부사격 조사도 온다). 기준 없음. 값: 순서가 어긋나면 +1, 맞거나 뒤가 없으면 -1
package rule

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

func ChainOrder(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	p := PieceOf(ctx)
	if p.Next < 0 || (p.Slot == 2 && p.Next == 1) {
		return -1, nil
	}
	if p.Slot >= 3 && p.Next <= 4 {
		return 1, nil
	}
	return Applies(p.Next < p.Slot), nil
}
