//ff:func feature=token type=reader control=sequence
//ff:what edgeFacts — the facts of a stretch of one sign at a unit's edge, from start to stop in the unit's letters: how many stand in a row, what stands on each side in the unit, whether the stretch is the whole unit, whether it stands after the letters that are read, whether those letters begin right after it, whether it begins an input line, and whether a space stands right after it
package token

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
	return m
}
