//ff:func feature=token type=rule control=iteration dimension=1
//ff:what verdictOf — evaluate one graph for one piece: the claim's verdict, and the exceptions that applied
package token

import (
	"strings"

	"github.com/geul-org/koparse/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// word is the unit's text up to the piece's end, which a closed-set
// exception compares. A graph that fails gives -1.
func verdictOf(g *toulmin.Graph, p rule.Piece, word string) (float64, []string) {
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyPiece, p)
	ctx.Set(rule.KeyWord, &rule.Word{Text: word, Sent: &rule.Sentence{}})
	results, err := g.Evaluate(ctx, toulmin.EvalOption{Trace: true})
	if err != nil || len(results) == 0 {
		return -1, nil
	}
	var why []string
	for _, t := range results[0].Trace {
		name, _, _ := strings.Cut(t.Name, "#")
		if t.Role == "except" && t.RuleVerdict > -1 {
			why = append(why, reasons[name])
		}
	}
	return results[0].Verdict, why
}
