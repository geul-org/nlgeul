//ff:func feature=parse type=helper control=sequence
//ff:what nearIn — a candidate of a word as the rules of a sentence read it (nearOf), with whether the word is written onto the word before it (the node's Glued), whether no adnominal may stand right before it by any of its candidates (unmodifiable), and whether each of its candidates is an adnominal (adnominalAll)
package parse

import "github.com/geul-org/nlgeul/pkg/rule"

func nearIn(n Node, br Branch) rule.Near {
	near := nearOf(br)
	near.Glued = n.Contest.Glued
	near.Unmodifiable = unmodifiable(n)
	near.AdnominalAll = adnominalAll(n)
	return near
}
