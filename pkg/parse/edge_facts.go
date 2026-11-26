//ff:func feature=parse type=reader control=sequence
//ff:what edgeFacts — the facts of a stretch of one sign at a unit's edge, from start to stop in the unit's letters: how many stand in a row, what stands on each side in the unit, whether the stretch is the whole unit, whether it stands after the letters that are read, whether those letters begin right after it, whether it begins an input line, whether a space stands right after it, whether a punctuation mark or symbol stands right next to it (a sign of the unit table read with the number before it is none: the % of 90%+), whether a sign that is no punctuation mark of a sentence stands in its run of signs (drawn), and the facts the cutter told of the unit's place (its line holds no letter or digit, a sum is written there, and, told only of a stretch that is the whole unit, a space or the sentence's edge stands on both sides of it)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// after tells that the signs stand after the unit's letters, and last
// that the stretch is the last of the signs before them.
func (sc *scan) edgeFacts(u unit, start, stop int, after, last bool) rule.Mark {
	all := []rune(u.text)
	m := rule.Mark{Run: stop - start}
	if start > 0 {
		m.Left = sideOf(string(all[start-1]))
	}
	if stop < len(all) {
		m.Right = sideOf(string(all[stop]))
	}
	if (start > 0 && signRune(all[start-1]) && !unitSign(all, start-1)) || (stop < len(all) && signRune(all[stop])) {
		m.Facts |= rule.MarkBySign
	}
	if start == 0 && stop == len(all) {
		m.Facts |= rule.MarkAlone
	}
	if after && len(sc.s) > 0 {
		m.Facts |= rule.MarkEnds
	}
	if !after && last && len(sc.s) > 0 {
		m.Facts |= rule.MarkLeads
	}
	if u.first && start == 0 {
		m.Facts |= rule.MarkHeads
	}
	if u.spaced && stop == len(all) {
		m.Facts |= rule.MarkSpaced
	}
	if drawn(all, start, stop) {
		m.Facts |= rule.MarkDrawn
	}
	m.Facts |= u.facts
	if m.Facts&rule.MarkAlone == 0 {
		m.Facts &^= rule.MarkApart
	}
	return m
}
