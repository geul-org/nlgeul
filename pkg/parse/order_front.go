//ff:func feature=parse type=reader control=sequence
//ff:what orderFront — the prefix 제 right before a number, when its graph stands: the number then begins at 제 and is named with it too (차례); the letters before 제 count as nothing on its left when they end a number (numberBefore)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func (sc *scan) orderFront(start int, names []string) (int, []string) {
	if start < 1 || sc.s[start-1] != orderHead {
		return start, names
	}
	m := rule.Mark{Right: rule.SideAlnum | rule.SideDigit, Run: 1}
	if start >= 2 && !sc.numberBefore(start-1) {
		m.Left = sideOf(string(sc.s[start-2]))
	}
	candidates := sc.judgeMark([]*entry{sc.ix.order}, m, string(orderHead))
	if len(candidates) == 0 {
		return start, names
	}
	return start - 1, union(markNames(shownNames(candidates)), names)
}
