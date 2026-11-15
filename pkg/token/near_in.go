//ff:func feature=token type=helper control=sequence
//ff:what nearIn — a candidate of a word as the rules of a sentence read it (nearOf), with whether the word is written onto the word before it (the node's Glued)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func nearIn(n Node, br Branch) rule.Near {
	near := nearOf(br)
	near.Glued = n.Contest.Glued
	return near
}
