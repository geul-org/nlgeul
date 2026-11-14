//ff:type feature=token type=model
//ff:what judgeKey — what a graph was asked: the graph, every fact of the piece it is given, and the letters of the word it is given
package token

import (
	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

type judgeKey struct {
	g     *toulmin.Graph
	facts rule.Piece
	word  string
}
