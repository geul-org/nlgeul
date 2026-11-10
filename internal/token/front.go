//ff:func feature=token type=reader control=sequence
//ff:what front — the sign right before a number, when it is one that stands at a number's front and its graph stands: the number then begins at the sign and is named with it too
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) front(start int, names string) (int, string) {
	if start < 1 {
		return start, names
	}
	e, ok := sc.ix.fronts[sc.s[start-1]]
	if !ok {
		return start, names
	}
	m := rule.Mark{Right: rule.SideAlnum | rule.SideDigit, Run: 1}
	if start >= 2 {
		m.Left = sideOf(string(sc.s[start-2]))
	}
	joined, _ := sc.judgeMark([]*entry{e}, m, string(sc.s[start-1]))
	if joined == "" {
		return start, names
	}
	return start - 1, union(shownNames(joined), names)
}
