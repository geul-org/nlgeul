//ff:func feature=rule type=helper control=sequence
//ff:what TestSlotBefore — a lower slot comes before a higher one and before a closing; 었 may come before 었; nothing else repeats or goes back
package rule

import "testing"

func TestSlotBefore(t *testing.T) {
	if !SlotBefore(SlotHonor, SlotPast) || !SlotBefore(SlotRecall, SlotClosing) || !SlotBefore(SlotPast, SlotPast) || SlotBefore(SlotGuess, SlotPast) || SlotBefore(SlotGuess, SlotGuess) {
		t.Fatal("SlotBefore")
	}
}
