//ff:func feature=parse type=reader control=sequence
//ff:what matchJamo — read a closing that puts ㄴ, ㄹ or ㅂ under the syllable before it (간다, 할까, 갑니다, 입니다)
package parse

import "github.com/geul-org/koparse/pkg/rule"

// A final consonant alone (no syllable after it) is read only where
// bareBase allows.
// Under 이 with a word before it the copula stem is read, and the
// whole syllable is part of the tail.
func matchJamo(s []rune, end int, e ending) (tailFact, bool) {
	start, ok := suffixAt(s, end, e.form)
	if !ok || start < 1 || rule.JongOf(s[start-1]) != e.lead {
		return tailFact{}, false
	}
	base := s[start-1] - rune(e.lead)
	if e.form == "" && !bareBase(s[:start], e.lead) {
		return tailFact{}, false
	}
	name := rule.LeadText(e.lead) + e.form
	rest, listed := closingPieces[name]
	if !listed {
		rest = e.form
	}
	t := tailFact{name, e.kind, string(s[:start-1]) + string(base), 2*(end-start) + 1, false, e.count, affixMark(s[:start], true) + rest}
	if base == '이' && start >= 2 && e.flags&fNoIda == 0 {
		t.copula, t.count, t.size = true, idaCount, t.size+1
		t.marked = affixMark(s[:start-1], false) + "+" + string(s[start-1]) + rest
	}
	return t, true
}
