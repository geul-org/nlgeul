//ff:func feature=parse type=reader control=sequence
//ff:what finish — after a closing is read at start: read the slots before it, require a stem, and report the tail as a fact with the word divided into its pieces
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

// closing is the closing as it is printed, with its leading "+". A
// closing that attaches only to 이 or 아니 (에요, 야) is not read
// anywhere else.
func finish(s []rune, start int, e ending, size int, closing string) (tailFact, bool) {
	i, size, fused := prefinals(s, start, size)
	if i < 1 {
		return tailFact{}, false
	}
	copula, extra := readsCopula(s, i, e)
	t := tailFact{rule.LeadText(e.lead) + e.form, e.kind, stemOf(s[:i]), size + extra, copula, e.count, ""}
	if e.flags&fNeedIda != 0 && !copula && t.stem != "아니" {
		return tailFact{}, false
	}
	if copula {
		t.count = idaCount
	}
	t.marked = stemMark(s, i, fused, extra > 0) + slotsMark(s[i:start]) + closing
	return t, true
}
