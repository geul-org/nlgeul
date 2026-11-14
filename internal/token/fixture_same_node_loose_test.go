//ff:func feature=token type=helper control=iteration dimension=1
//ff:what sameNodeLoose — are two nodes the same but for the numbers of birth: the span, unit, Glued and letters, and every candidate with its words, verdict and preference, in order (the order of the candidates holds what the numbers of birth say among the same lambda, and the numbers themselves are those of another count)
package token

import "reflect"

func sameNodeLoose(a, b Node) bool {
	ca, cb := a.Contest, b.Contest
	if a.Span != b.Span || ca.Unit != cb.Unit || ca.Glued != cb.Glued || ca.Text != cb.Text || len(ca.Branches) != len(cb.Branches) {
		return false
	}
	for k := range ca.Branches {
		x, y := ca.Branches[k], cb.Branches[k]
		if !reflect.DeepEqual(x.Words, y.Words) || x.Verdict != y.Verdict || x.Pref != y.Pref {
			return false
		}
	}
	return true
}
