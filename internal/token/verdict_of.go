//ff:func feature=token type=rule control=sequence
//ff:what verdictOf — evaluate one graph for one piece: the claim's verdict, and the exceptions that applied
package token

import (
	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// word is the unit's text up to the piece's end, which a closed-set
// exception compares.
func verdictOf(g *toulmin.Graph, p rule.Piece, word string) (float64, []string) {
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyPiece, p)
	ctx.Set(rule.KeyWord, &rule.Word{Text: word, Sent: &rule.Sentence{}})
	return evaluate(g, ctx)
}
