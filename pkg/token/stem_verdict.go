//ff:func feature=token type=rule control=sequence
//ff:what stemVerdict — the verdict of a kind's graph for a stem with the facts given (what follows it, its observed count, its letters, the word before it); the same facts are judged once and remembered
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func (ix *index) stemVerdict(kind int, p rule.Piece) float64 {
	return ix.judged(ix.kinds[kind], p, "")
}
