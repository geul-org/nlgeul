//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameNode — are two nodes the same: the span, unit, Glued and letters, and every candidate with its words (as they are), verdict, preference and number of birth, in order
package token

import "reflect"

func sameNode(a, b Node) bool {
	ca, cb := a.Contest, b.Contest
	if a.Span != b.Span || ca.Unit != cb.Unit || ca.Glued != cb.Glued || ca.Text != cb.Text || len(ca.Branches) != len(cb.Branches) {
		return false
	}
	for k := range ca.Branches {
		x, y := ca.Branches[k], cb.Branches[k]
		if !reflect.DeepEqual(x.Words, y.Words) || x.Verdict != y.Verdict || x.Pref != y.Pref || bornOf(x) != bornOf(y) {
			return false
		}
	}
	return true
}
