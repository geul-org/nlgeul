//ff:func feature=token type=helper control=sequence
//ff:what graphHolds — whether the claim of a graph of a rule of a sentence stands for a candidate with the words around it (rule.Around), evaluated with no trace
package token

import (
	"testing"

	"github.com/geul-org/nlgeul/pkg/rule"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func graphHolds(t *testing.T, r sentenceRule, a rule.Around) bool {
	t.Helper()
	ctx := toulmin.NewContext()
	ctx.Set(rule.KeyAround, a)
	v, _ := evaluate(r.graph, ctx, false)
	return v > 0
}
