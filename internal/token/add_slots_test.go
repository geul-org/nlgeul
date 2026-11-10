//ff:func feature=token type=reader control=sequence
//ff:what TestAddSlots — the slots are indexed with their order; the fused ㅆ under its key
package token

import "testing"

func TestAddSlots(t *testing.T) {
	ix := &index{byKey: map[rune][]*entry{}, words: map[string]*entry{}, closed: map[string]int{}}
	ix.addSlots()
	if ix.byKey['었'][0].slot != 2 || ix.byKey['었'][0].class != classSlot || ix.byKey['ㅆ'][0].lead != jongSS || ix.byKey['더'][0].slot != 4 || ix.graphs != 7 {
		t.Fatal(ix.graphs)
	}
}
