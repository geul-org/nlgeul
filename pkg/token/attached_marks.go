//ff:func feature=token type=helper control=sequence
//ff:what attachedMarks — does a candidate read its word as markers alone that attach to the word before: no stem, and its first marker is of a kind of apartEnds (a particle, the copula 이, a closing), not a punctuation mark
package token

import (
	"slices"

	"github.com/geul-org/nlgeul/pkg/rule"
)

func attachedMarks(n rule.Near) bool {
	marks := rule.MarksOf(n)
	if n.Stem != "" || n.Kind != "" || len(marks) == 0 {
		return false
	}
	return slices.Contains(apartEnds, labelOf(marks[0]))
}
