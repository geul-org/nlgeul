//ff:func feature=rule type=helper control=sequence
//ff:what SlotBefore — 조건식(규칙 SlotOrder와 읽기 코드가 함께 쓴다). 칸 slot이 칸 next 바로 앞에 올 수 있는가: 순서가 오르거나, 었이 겹친다
package rule

func SlotBefore(slot, next int) bool {
	return slot < next || (slot == next && slot == SlotPast)
}
