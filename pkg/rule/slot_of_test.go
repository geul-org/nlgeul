//ff:func feature=rule type=helper control=sequence
//ff:what TestSlotOf — 시, 었 았 였, 겠, 더 are slots in that order; another syllable is none
package rule

import "testing"

func TestSlotOf(t *testing.T) {
	if SlotOf('시') != SlotHonor || SlotOf('었') != SlotPast || SlotOf('았') != SlotPast || SlotOf('였') != SlotPast || SlotOf('겠') != SlotGuess || SlotOf('더') != SlotRecall || SlotOf('가') != 0 {
		t.Fatal("SlotOf")
	}
}
