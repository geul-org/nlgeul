//ff:func feature=parse type=reader control=iteration dimension=1
//ff:what prefinals — read the slots left of a closing in their order 시 < 었 < 겠 < 더 (endings.md section 1); where the stem ends, the length read, and whether a ㅆ fused into the stem was read
package parse

import "github.com/geul-org/koparse/pkg/rule"

// Every slot read adds a syllable (2). 더, 겠 and 었/았/였 are read
// from the right while their order holds, 었/았/였 twice at most. A
// fused ㅆ belongs to the stem's last syllable: it adds 1 and reading
// stops there. 시 is read last.
func prefinals(s []rune, start, size int) (int, int, bool) {
	i, next, pasts := start, rule.SlotClosing, 0
	for i > 0 {
		slot := rule.SlotOf(s[i-1])
		if slot < rule.SlotPast || !rule.SlotBefore(slot, next) || (slot == rule.SlotPast && pasts == 2) {
			break
		}
		if slot == rule.SlotPast {
			pasts++
		}
		i, size, next = i-1, size+2, slot
	}
	if i > 0 && fusedPast(s[i-1]) {
		return i, size + 1, true
	}
	if i > 0 && rule.SlotOf(s[i-1]) == rule.SlotHonor {
		i, size = i-1, size+2
	}
	return i, size, false
}
