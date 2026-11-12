//ff:func feature=token type=rule control=sequence
//ff:what fact — the criterion of rule.PieceIs for some yes/no facts of a piece (rule.Piece…)
package token

import "github.com/geul-org/nlgeul/pkg/rule"

func fact(mask int) *rule.Bits {
	return &rule.Bits{Of: "piece", Mask: mask}
}
