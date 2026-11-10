//ff:func feature=token type=rule control=sequence
//ff:what newWord — the verdict of the new word's graph for a token sequence no table lists: the usual low one, or none when the sequence stands after a boundary and reads as markers
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (ix *index) newWord(marked bool) float64 {
	if !marked {
		return ix.unknown
	}
	v, _ := verdictOf(ix.stem, rule.Piece{Marked: true}, "")
	return v
}
