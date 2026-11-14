//ff:func feature=token type=helper control=iteration dimension=2
//ff:what wayKnown — whether a way is of known words only: no part of the first candidate of any of its words is a new word (Part.Known); every candidate of a node is one word
package token_test

import "github.com/geul-org/nlgeul/pkg/token"

func wayKnown(seg token.Segment) bool {
	for _, n := range seg.Nodes {
		for _, p := range n.Contest.Branches[0].Words[0].Parts {
			if !p.Known() {
				return false
			}
		}
	}
	return true
}
