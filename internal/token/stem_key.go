//ff:type feature=token type=model
//ff:what stemKey — what a kind's graph was asked: the kind and every fact of the stem the graph is given
package token

import "github.com/geul-org/nlgeul/pkg/rule"

type stemKey struct {
	kind  int
	facts rule.Piece
}
